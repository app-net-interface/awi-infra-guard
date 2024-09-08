# CI Pipeline
To streamline the works on the project, CI pipeline is available for developers. It uses GitHub Actions, Terraform and grpc_cli among others and allows to run all the tests locally using self-hosted runners.

The pipeline consists of 3 workflows, each specified under the [.github/workflows](.github/workflows) directory. Each of the workflows provides a separate functionality - resource creation, testing and deletion.

## Self-hosted runner
A self-hosted runner is required to run any workflow. To create one locally, please follow the [GitHub's documentation](https://docs.github.com/en/actions/hosting-your-own-runners/managing-self-hosted-runners/adding-self-hosted-runners). A label should also be added to the runner. It will be later used to specify a runner on which to execute workflows. Example label:
``
<github_handle>-<runner_name>
``

Additionally a GitHub CLI tool (``gh``) is required to execute workflows from a local machine. Please follow [GitHub's documentation](https://github.com/cli/cli#installation) to install it.

## Workflow execution
Once runner is crated and running, a workflow can be executed with the following command
```
gh workflow run <workflow_filename> --ref <branch_name> -f self_hosted_runner_name=<self_hosted_runner_label>
```
``workflow_filename`` - file containing instructions to execute using GitHub Actions</br>
``branch_name`` - a branch on which to execute a workflow</br>
``self_hosted_runner_label`` - previously specified label, allows to choose a runner on which the workflow will be executed 

## AWS access
Please specify access credentials to an AWS account using [GitHub Secrets](https://docs.github.com/en/actions/security-for-github-actions/security-guides/using-secrets-in-github-actions#creating-secrets-for-a-repository). Two secrets are required: ``AWS_ACCESS_KEY_ID`` and ``AWS_SECRET_ACCESS_KEY``. Please prepend both of these names with an appropriate github handle ``<GITHUB_HANDLE_>``. </br>
There can be more credentials defined in a repository, to allow usage of multiple AWS accounts the prepended github handle will be used to automatically select credentials based on the user that triggered the workflow.

The following secrets should be specified before any workflow is executed:
```
<GITHUB_HANDLE>_AWS_ACCESS_KEY_ID
<GITHUB_HANDLE>_AWS_SECRET_ACCESS_KEY
```

## Create workflow
It uses Terraform to spin up a sample infrastructure on AWS. Created resources can be later used for testing purposes.

Please run the following command to execute the create workflow:
```
gh workflow run create-aws-test-infrastructure.yml --ref <branch_name> -f self_hosted_runner_name=<self_hosted_runner_label>
```

Please notice that the .tfstate file, created during the workflow run, is kept locally under the ``terraform/aws/backend`` directory in a repository that lives where the self-hosted runner is located, e.g. ``~/actions-runner/_work/awi-infra-guard/awi-infra-guard/terraform/aws/backend``

## Test workflow
Runs checks on the created infrastructure, ensuring that the required fields are present and accessible by the awi-infra-guard.

An instance of the awi-infra-guard, accepting the grpc calls, must be running locally, on the same machine as the self-hosted runner. The tests will be executed against that instance. Please refer to [README.md](README.md) for more information on running the app.

Please run the following command to execute the test workflow:
```
gh workflow run test-aws-infrastructure.yml --ref <branch_name> -f self_hosted_runner_name=<self_hosted_runner_label>
```

## Destroy workflow
Deletes the sample infrastrcture created by the [Create workflow](#create-workflow)

Please run the following command to execute the destroy workflow:
```
gh workflow run destroy-aws-test-infrastructure.yml --ref <branch_name> -f self_hosted_runner_name=<self_hosted_runner_label>
```

## Workflow results
All the results and logs from the workflow runs are available on the GitHub page of the repository, under the Actions tab. 

Optionally they can accessed using the following command:
```
gh run view
```