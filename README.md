# Incept - Terraform for AI

## Design
Users declare what they want
```toml
[params]
train_steps = 1000
batch_size = 32

[[services.train]]
name = "train"
dataset = "mnist@train"
model = "multilayer perceptron"
ckpt = "./ckpt/model"
batch_size = "{{ batch_size }}"
steps = "{{ train_steps }}"

[[services.eval]]
dependencies = [ "train" ]
dataset = "mnist"
dataset_split_rng = "{{ dataset_rng }}"
dataset_split_ratio = 80
ckpt = "./ckpt/model"
batch_size = 32
```

Then incept will figure out what to do and do things correctly.
