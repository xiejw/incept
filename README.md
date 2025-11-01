# Incept - Terraform for AI - Simple, Declarative, Reproducible.

## Design
Users declare what they want
```toml
[params]
rng_seed = 123
train_steps = 1000
batch_size = 32

[[services.train]]
name = "train"
dataset = "mnist@train"
model = "mlp@v1"
ckpt = "./ckpt/model"
batch_size = "{{ batch_size }}"
steps = "{{ train_steps }}"
rng_seed = "{{ rng_seed }}"

[[services.eval]]
dependencies = [ "train" ]
dataset = "mnist@eval"
ckpt = "./ckpt/model"
batch_size = 32
```

Then incept will figure out what to do and do things correctly.
