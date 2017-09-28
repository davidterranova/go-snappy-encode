// Created by davidterranova on 27/09/2017.

package snappyencode

import (
	"encoding/base64"
	"testing"
)

type Dumb struct {
	Text string
	Num  int
}

func (d *Dumb) equal(dd *Dumb) bool {
	if d.Text == dd.Text && d.Num == dd.Num {
		return true
	} else {
		return false
	}
}

func TestEncode(t *testing.T) {
	var o Dumb = Dumb{
		"plop",
		10,
	}
	enc, err := Encode(o)
	if err != nil {
		t.Error(err)
	}
	if len(enc) == 0 {
		t.Error("encoded data should not be empty")
	}
}

func TestDecode(t *testing.T) {
	var stringB64 = "/wYAAHNOYVBwWQE0AADacdMmI/+BAwEBBER1bWIB/4IAAQIBBFRleHQBDAABA051bQEEAAAAC/+CAQRwbG9wARQA"
	enc, err := base64.StdEncoding.DecodeString(stringB64)
	if err != nil {
		t.Error("impossible to decode b64 string")
	}

	var o Dumb
	err = Decode(enc, &o)
	if err != nil {
		t.Error(err)
	}

	if o.Text != "plop" || o.Num != 10 {
		t.Error("Object not well decoded")
	}
}

func TestEncDec(t *testing.T) {
	var o Dumb = Dumb{
		"plop",
		10,
	}
	enc, err := Encode(o)
	if err != nil {
		t.Error(err)
	}
	if len(enc) == 0 {
		t.Error("encoded data should not be empty")
	}

	var oo Dumb
	err = Decode(enc, &oo)
	if err != nil {
		t.Error(err)
	}

	if !o.equal(&oo) {
		t.Error("bad encode / decode, objects !=")
	}
}
