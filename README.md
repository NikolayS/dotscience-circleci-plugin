# Dotscience CircleCI plugin

This plugin can be used in Dotscience pipeline step to start a new CircleCI job build.

## Usage

Create .dotscience.yml either in your Github repo (if you are using clone & run) or just workspace root.

Example configuration:

```yaml
kind: pipeline

after:
- name: circleci
  image: quay.io/dotmesh/dotscience-circleci-plugin
  settings:
    # best to use project API token
    # info here: https://circleci.com/docs/2.0/managing-api-tokens/#creating-a-project-api-token
    token: xxx
    # your CircleCI username
    username: your-username
    # your CircleCI project name
    project: dotscience-pipeline-demo
```

Then, run the task (assuming you have a `train.py` file there to train your model):

```
ds run -v -p pipeline-test --upload-path . python train.py
```