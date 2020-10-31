extern crate openssl;
extern crate base64;

use openssl::rsa::{Rsa, Padding};
use base64::{CharacterSet, Config};
use string_builder::Builder;


fn decrypt(encrypted: Vec<u8>, pubkey_pem: String) -> String {
    println!("Encrypted len: {}", encrypted.len());
    let rsa = Rsa::public_key_from_pem(pubkey_pem.as_bytes()).unwrap();
    let mut builder = Builder::default();
    let mut i = 0;
    while i < encrypted.len() {
        let mut buf: Vec<u8> = vec![0; rsa.size() as usize];
        let _ = rsa.public_decrypt(&encrypted[i..i+128], &mut buf, Padding::PKCS1).unwrap();
        builder.append(String::from_utf8(buf).unwrap());
        i += 128;
    }
    return builder.string().unwrap();
}

fn main() {
    let public_key_pem = "-----BEGIN PUBLIC KEY-----\nMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDIqgPyECauK+RWC2H+BK4Mi1owDDpnwzeSEjnKiP8y/jNhXM9QlcTv7x49NvdOF4U1F/GBFc7+vfkvSXldv+Qh1KA+hztDdy3Rh4sXJSB9Dsu6zuas3qDqOY34xQgdcqI2P0xk9JRyUzGQnWHXxQtMC8khuZRroT12dHn6xEu5qwIDAQAB\n-----END PUBLIC KEY-----";
    let data = "FPvxYW7YFGeOTOxCZKrjg6/gmsr2hO8nOYT8hIAh5Vk/WiGU9zTTenly5id1JwLNQFi6HswvxFkyXukUk1dzg+VXSNjpXFbj+5/0Kv6eMzB6+h0mOt7oPMc/mORbigZHaL6FdgKCbmynYcx2ehp37WUg0MfMYHNlMS3yu3Fe8wBLLR0gKUIKJpv7pStcoCMXGSWNrkDRL1olYLDGeceTMAwt5g8VYz9loP55QrE4I2b0Hw1dxJpjX6ngUnyi3NWcD41DNe+aVUUtlxXy2Ptg0bnoy85trnNnpw3hURm3vUlm4E8lYmdRu/WLbJWh7dCY8bkQOH5FyxvgpyhGVIjFvw==";
    let decoded = base64::decode_config(data, Config::new(CharacterSet::Standard,  true)).unwrap();

    let res = decrypt(decoded, public_key_pem.to_string());

    println!("Decrypted: {}", res);
}
