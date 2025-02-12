# Gatekeeper Policy development and testing

## What's Gatekeeper
official doc https://open-policy-agent.github.io/gatekeeper/website/docs/

## Guardrails folder structures

There are several folders under guardrails:

* gktemplates - the Constraint Templates used by gatekeeper, which are generated through generate.sh, do *not* modify them.
* gkconstraints - the Constraints that are used by gatekeeper together with Constraint Templates.

* gktemplates-src - the rego src file for Constraint Templates, consumed by generate.sh
* scripts - generate.sh will combine src.rego and *.tmpl to form actual Constraint Templates under gktemplates. test.sh executes the rego tests under each gktemplates-src subfolder.
* staticresources - yaml resources for gatekeeper deployment

## Policy structure
Each policy contains 2 parts, [ConstraintTemplate](https://open-policy-agent.github.io/gatekeeper/website/docs/constrainttemplates/) and [Constraint](https://open-policy-agent.github.io/gatekeeper/website/docs/howto/#constraints)

ConstraintTemplate, ie,. gktemplate/\$TEMPLATE_NAME.yaml, is generated by policies/scripts/generate.sh, who combines the yaml part, ie,. the gktemplates-src/\$TEMPLATE_NAME/\$TEMPLATE_NAME.tmpl and rego program (gktemplates-src/\$TEMPLATE_NAME/src.rego).
please dont edit the ConstraintTemplate directly, just provide the corresponding $TEMPLATE_NAME.tmpl and src.rego, the generate.sh will produce the ConstraintTemplate file.

Constraint is manually created

## Create new policy

* Create a new subfolder for each new Constraint Template under gktemplates-src
* Create a tmpl file with unique and meaningful name in above subfolder, which contains everything except for the rego, example:

```yaml
apiVersion: templates.gatekeeper.sh/v1
kind: ConstraintTemplate
metadata:
  name: arodenyprivilegednamespace
  annotations:
    metadata.gatekeeper.sh/title: "Privileged Namespace"
    metadata.gatekeeper.sh/version: 1.0.0
    description: >-
      Disallows creating, updating or deleting resources in privileged namespaces.
spec:
  crd:
    spec:
      names:
        kind: ARODenyPrivilegedNamespace
      validation:
        # Schema for the `parameters` field
        openAPIV3Schema:
          type: object
          description: >-
            Disallows creating, updating or deleting resources in privileged namespaces.
  targets:
    - target: admission.k8s.gatekeeper.sh
      rego: |
{{ file.Read "gktemplates-src/aro-deny-privileged-namespace/src.rego" | strings.Indent 8 | strings.TrimSuffix "\n" }}
      libs:
        - |
{{ file.Read "gktemplates-src/library/common.rego" | strings.Indent 10 | strings.TrimSuffix "\n" }}

```


* Create the src.rego file in the same folder, howto https://www.openpolicyagent.org/docs/latest/policy-language/, example:
```
package arodenyprivilegednamespace

import data.lib.common.is_priv_namespace
import data.lib.common.is_exempted_account
import data.lib.common.get_username

violation[{"msg": msg}] {
  ns := input.review.object.metadata.namespace
  is_priv_namespace(ns)
  not is_exempted_account(input.review)
  username := get_username(input.review)
  msg := sprintf("user %v not allowed to operate in namespace %v", [username, ns])
}
```
* Create src_test.rego for unit tests in the same foler, which will be called by test.sh, howto https://www.openpolicyagent.org/docs/latest/policy-testing/, example:
```
package arodenyprivilegednamespace

test_input_allowed_ns {
  input := { "review": input_ns(input_allowed_ns) }
  results := violation with input as input
  count(results) == 0
}

test_input_disallowed_ns1 {
  input := { "review": input_ns(input_disallowed_ns1) }
  results := violation with input as input
  count(results) == 1
}

input_ns(ns) = output {
  output = {
    "object": {
      "metadata": {
        "namespace": ns
      }
    }
  }
}

input_allowed_ns = "mytest"

input_disallowed_ns1 = "openshift-config"
```

* Create [Constraint](https://open-policy-agent.github.io/gatekeeper/website/docs/howto/#constraints) for the policy, example:

```yaml
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: ARODenyPrivilegedNamespace
metadata:
  name: aro-privileged-namespace-deny
spec:
  enforcementAction: {{.Enforcement}}
  match:
    kinds:
      - apiGroups: [""]
        kinds: [
        "Pod",
        "Secret",
        "Service",
        "ServiceAccount",
        "ReplicationController",
        "ResourceQuota",
        ]
      - apiGroups: ["apps"]
        kinds: ["Deployment", "ReplicaSet", "StatefulSet", "DaemonSet"]
      - apiGroups: ["batch"]
        kinds: ["Job", "CronJob"]
      - apiGroups: ["rbac.authorization.k8s.io"]
        kinds: ["Role", "RoleBinding"]
      - apiGroups: ["policy"]
        kinds: ["PodDisruptionBudget"]
```

Make sure the filename of constraint is the same as the .metadata.name of the Constraint object, as it is the feature flag name that will be used to turn on / off the policy.

## Test Rego source code

* install opa cli, refer https://github.com/open-policy-agent/opa/releases/

* after _test.go is done, test it out, and fix the problem
  ```sh
  opa test ../library/common.rego *.rego [-v] #-v for verbose
  ```

## Generate the Constraint Templates

* install gomplate which is used by generate.sh, see https://docs.gomplate.ca/installing/

* execute generate.sh under policies, which will generate the acutal Constraint Templates to gktemplates folder, example:


  ```sh
  # Generate all the Constraint Templates
  ARO-RP/pkg/operator/controllers/guardrails/policies$ ./scripts/generate.sh 
  Generating gktemplates/aro-deny-delete.yaml from gktemplates-src/aro-deny-delete/aro-deny-delete.tmpl
  Generating gktemplates/aro-deny-privileged-namespace.yaml from gktemplates-src/aro-deny-privileged-namespace/aro-deny-privileged-namespace.tmpl
  Generating gktemplates/aro-deny-labels.yaml from gktemplates-src/aro-deny-labels/aro-deny-labels.tmpl
  ```

  ```sh
  # Generate a specific Constraint Template by providing the specific policy directory under gktemplates-src folder
  ARO-RP/pkg/operator/controllers/guardrails/policies$ ./scripts/generate.sh aro-deny-machine-config
  Generating gktemplates/aro-deny-machine-config.yaml from gktemplates-src/aro-deny-machine-config/aro-deny-machine-config.tmpl
  ```

## Test policy with gator

Create suite.yaml and testcases in gator-test folder under the folder created for the new policy, refer example below:

```yaml
kind: Suite
apiVersion: test.gatekeeper.sh/v1alpha1
metadata:
  name: privileged-namespace
tests:
- name: privileged-namespace
  template: ../../gktemplates/aro-deny-privileged-namespace.yaml
  constraint: ../../gkconstraints-test/aro-privileged-namespace-deny.yaml
  cases:
  - name: ns-allowed-pod
    object: gator-test/ns_allowed_pod.yaml
    assertions:
    - violations: no
  - name: ns-disallowed-pod
    object: gator-test/ns_disallowed_pod.yaml
    assertions:
    - violations: yes
      message: User test-user not allowed to operate in namespace openshift-config
  - name: ns-disallowed-deploy
    object: gator-test/ns_disallowed_deploy.yaml
    assertions:
    - violations: yes
      message: User test-user not allowed to operate in namespace openshift-config
```
gkconstraints-test here stores the target yaml files after expanding "{{.Enforcement}}" symbol.

gator tests ConstraintTemplate and Constraint together, items under cases keyword are test cases indicator, everyone pointing to a yaml file in gator-test, which provides test input for one scenario, example:

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: allowed
  namespace: test
spec:
  serviceAccountName: test-user
  containers:
    - name: test
      image: openpolicyagent/opa:0.9.2
      args:
        - "run"
        - "--server"
        - "--addr=localhost:8080"
      resources:
        limits:
          cpu: "100m"
          memory: "30Mi"
```
the `assertions` section is the expected result

gator test can be done via cmd:

test.sh executes both opa test and gator verify
```sh
# Run tests for all the policies
ARO-RP/pkg/operator/controllers/guardrails/policies$ ./scripts/test.sh
```
```sh
# Run tests for a specific policy
# Providing the policy directory under gktemplates-src folder and the correspondent Constraint file under gkconstraints folder
ARO-RP/pkg/operator/controllers/guardrails/policies$ ./scripts/test.sh aro-deny-machine-config aro-machine-config-deny.yaml
```

or below cmd after test.sh has been executed:
```sh
gator verify . [-v] #-v for verbose
```

Sometimes we need to mock kube admission review request especially as gator test inputs when verifying policies that check specific operations (e.g., CREATE, DELETE or UPDATE).

Please refer the yaml file below as a sample of kube admission review request:

```yaml
kind: AdmissionReview
apiVersion: admission.k8s.io/v1
request:
  uid: d700ab7f-8f42-45ff-83f5-782c739806d9
  operation: UPDATE
  userInfo:
    username: kube-review
    uid: 45884572-1cab-49e5-be4c-1d2eb0299776
  object:
    kind: MachineConfig
    apiVersion: machineconfiguration.openshift.io/v1
    metadata:
      name: 99-worker-generated-crio-fake
  oldObject:
    kind: MachineConfig
    apiVersion: machineconfiguration.openshift.io/v1
    metadata:
      name: 99-worker-generated-crio-seccomp-use-default
  dryRun: true
```
A tool [admr-gen](https://github.com/ArrisLee/admr-gen) has been created and can be utilized to generate mocked kube admission review requests in an easier way.

## Enable and test your policy on a dev cluster

Set up local dev env following “Deploy development RP” section if not already: https://github.com/Azure/ARO-RP/blob/master/docs/deploy-development-rp.md

Deploy a dev cluster $CLUSTER in your preferred region, cmd example:
```sh
CLUSTER=jeff-test-aro go run ./hack/cluster create
```

Scale the standard aro operator to 0, cmd:
```sh
oc scale -n openshift-azure-operator deployment/aro-operator-master --replicas=0
```

Run aro operator from local code, cmd example:
```sh
CLUSTER=jeff-test-aro go run -tags aro,containers_image_openpgp ./cmd/aro operator master
```

Wait a couple of minutes until aro operator fully synchronized

Enable guardrails, set it to managed, cmd:
```sh
oc patch cluster.aro.openshift.io cluster --type json -p '[{ "op": "replace", "path": "/spec/operatorflags/aro.guardrails.deploy.managed", "value":"true" }]'
oc patch cluster.aro.openshift.io cluster --type json -p '[{ "op": "replace", "path": "/spec/operatorflags/aro.guardrails.enabled", "value":"true" }]'
```
The sequence for above cmds is essential, please dont change the order!

Use below cmd to verify the gatekeeper is deployed and ready
```sh
$ oc get all -n openshift-azure-guardrails
NAME                                                READY   STATUS    RESTARTS   AGE
pod/gatekeeper-audit-67c4c669c7-mrr6w               1/1     Running   0          10h
pod/gatekeeper-controller-manager-b887b69bd-mzhsh   1/1     Running   0          10h
pod/gatekeeper-controller-manager-b887b69bd-tb8zc   1/1     Running   0          10h
pod/gatekeeper-controller-manager-b887b69bd-xnvv4   1/1     Running   0          10h

NAME                                 TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)   AGE
service/gatekeeper-webhook-service   ClusterIP   172.30.51.233   <none>        443/TCP   35h

NAME                                            READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/gatekeeper-audit                1/1     1            1           10h
deployment.apps/gatekeeper-controller-manager   3/3     3            3           10h

NAME                                                      DESIRED   CURRENT   READY   AGE
replicaset.apps/gatekeeper-audit-67c4c669c7               1         1         1       10h
replicaset.apps/gatekeeper-controller-manager-b887b69bd   3         3         3       10h
```

Verify ConstraintTemplate is created
```sh
$ oc get constrainttemplate
NAME            AGE
arodenylabels   20h
```


Enforce the machine rule, cmd:
```sh
oc patch cluster.aro.openshift.io cluster --type json -p '[{ "op": "replace", "path": "/spec/operatorflags/aro.guardrails.policies.aro-machines-deny.managed", "value":"true" }]'
oc patch cluster.aro.openshift.io cluster --type json -p '[{ "op": "replace", "path": "/spec/operatorflags/aro.guardrails.policies.aro-machines-deny.enforcement", "value":"deny" }]'
```
Note: the feature flag name is the corresponding Constraint FILE name, which can be found under pkg/operator/controllers/guardrails/policies/gkconstraints/, Eg, aro-machines-deny.yaml

Verify corresponding gatekeeper Constraint has been created:
```sh
$ oc get constraint
NAME                ENFORCEMENT-ACTION   TOTAL-VIOLATIONS
aro-machines-deny   deny
```

Once the constraint is created, you are all good to rock with your policy!