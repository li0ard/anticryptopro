from pygost.wrap import unwrap_cryptopro

KEY = "c2bb428f981f30d8e623db786d9bd780e7d9f785a950211530142ae46fab8dce"

ukm = "FEF0268E68E6A7EA"
cek_enc = "7987275594A59F491E62F636F794BDD49A5CD547E0D901DF3AD58DA9EFFC2657"
cek_mac = "68369E8A"

print("UNWRAP = " + unwrap_cryptopro(
	bytes.fromhex(KEY),
	bytes.fromhex(ukm + cek_enc + cek_mac)
).hex())
