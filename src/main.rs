use serde::Deserialize;
use std::collections::HashMap;
use std::fs;

#[derive(Deserialize, Debug)]
#[serde(tag = "model")]
pub enum Model {
    MLP {
        num_layers: usize,
        batch_size: usize,
    },
}

#[derive(Deserialize, Debug)]
#[serde(tag = "stage")]
pub enum Service {
    Train {
        dataset: String,
        model: Model,
        dependencies: Option<Vec<String>>,
    },
    Eval {
        dataset: String,
        model: Model,
        dependencies: Option<Vec<String>>,
    },
}

#[derive(Deserialize, Debug)]
pub struct Config {
    pub params: HashMap<String, toml::Value>,
    pub services: HashMap<String, Service>,
}

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let toml_str = fs::read_to_string("example.toml")?;
    let value: Config = toml::from_str(&toml_str).unwrap();
    println!("{:#?}", value);
    Ok(())
}
