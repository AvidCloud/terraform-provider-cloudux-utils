// Copyright 2022 by Avid Technology, Inc.
package provider

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

func resourceSiteKey() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSiteKeyCreate,
		ReadContext:   resourceSiteKeyRead,
		UpdateContext: resourceSiteKeyUpdate,
		DeleteContext: resourceSiteKeyDelete,
		Schema: map[string]*schema.Schema{
			"rsa_bits": &schema.Schema{
				Type:      schema.TypeInt,
				Required:  false,
				Default:   4096,
				Sensitive: false,
				ForceNew:  true,
			},
			"public_key": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"private_key": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func toPrivateKey(key *rsa.PrivateKey) (string, error) {
	data, err := x509.MarshalPKCS8PrivateKey(key)
	if err != nil {
		return "", err
	}
	var privateKey = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: data,
	}

	buf := new(bytes.Buffer)
	err = pem.Encode(buf, privateKey)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func toPublicKey(key *rsa.PrivateKey) (string, error) {
	data, err := x509.MarshalPKIXPublicKey(&key.PublicKey)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	err = pem.Encode(buf, &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: data,
	})
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func resourceSiteKeyCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	bitSize := d.Get("rsa_bits").(int)

	reader := rand.Reader
	key, err := rsa.GenerateKey(reader, bitSize)
	if err != nil {
		return diag.FromErr(err)
	}
	privateKey, err := toPrivateKey(key)
	if err != nil {
		return diag.FromErr(err)
	}
	publicKey, err := toPublicKey(key)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(hashForState("sitekey_" + strconv.Itoa(bitSize)))
	d.Set("public_key", publicKey)
	d.Set("private_key", privateKey)
	return diags
}

func resourceSiteKeyRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func resourceSiteKeyUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func resourceSiteKeyDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	d.SetId("")
	return nil
}
