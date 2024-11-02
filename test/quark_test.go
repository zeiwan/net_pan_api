package test

import (
	"fmt"
	"github.com/zeiwan/net_pan_api/base"
	"github.com/zeiwan/net_pan_api/module"
	"testing"
)

var q = base.NewCloudQuark()

const cookie = "ctoken=nMbRshRXN0trRjrX0xhu-97V; b-user-id=a2c9ea31-e534-0d8a-675b-160af4b2fb2b; grey-id=ec950091-e39c-0fb4-ac36-75f62c830762; grey-id.sig=Z7Wf49PHyos32ntFZ5ZL4Amp_DyqcZ_BdH8IyHJUgbc; isQuark=true; isQuark.sig=hUgqObykqFom5Y09bll94T1sS9abT1X-4Df_lzgl8nM; __wpkreporterwid_=8c358d29-b6db-4019-28f9-76d0b9c21c91; _UP_A4A_11_=wb96a128ce55457ea2cc95945d87e4b6; _UP_D_=pc; _UP_30C_6A_=st96a620197nb6xgklsq2ryx895h3egw; _UP_TS_=sg1be132b1cfb78bef9af0eeb9bf0a0c25e; _UP_E37_B7_=sg1be132b1cfb78bef9af0eeb9bf0a0c25e; _UP_TG_=st96a620197nb6xgklsq2ryx895h3egw; _UP_335_2B_=1; __pus=c11e16d728126816a758207921c0248cAATsRU8Ae+xOOieWPH5HWahIpStieaZjLKct/x4E1pmg833DW7yizu3bNG9X6k47cpHJMxuKAxRUj5sOVgpqkRGn; __kp=a50d7ee0-9459-11ef-be06-fdb184d33df9; __kps=AATrbC2XxtcPZEXZtA22lSZZ; __ktd=uGRoTefKBPGQ1VcV7XIuLw==; __uid=AATrbC2XxtcPZEXZtA22lSZZ; tfstk=fUMEC6vdOppeGMJ5Fkyz05e9SZ2LUJYjtYa7q0muAy43AJio_0i2vugIqzlzS4KddJYdEzmxWWanPJ_yUVy3wXa7Rzyr2q-6lK9jvDe8KETXhzjf91ef-ab7Z5jiavqhvK9jvDcG-8ryhvGCbMz4r8V3ZO2gWk2lKW0uS5q0cw2or8xasoZ8r943-l4gDP4uE80ojjBqx0XacWxX0awJme8btrm3-DiiIx_brD4NET4UYWoZxPWlEAc_HRA_-II_yJaKfkuJH9e3ajn742JFK4oxBVrioLWb-cnjGWHHH_nLKyyiKjSlEcyU-SDiSw8aovniHRV6nnoiB2iKQ0skElgbSDH3UKxYLJ43QlHWW94rsbDQ98QHlWno_VrUng71ylxFPYhFr_V3XlzXbhlm4a3O-nak__C8s-EalHZCw_F3plzXbpCRwWd8brtpA; __itrace_wid=b9ccf2df-a115-4324-1bc6-87f4beb6458a; __puus=2e0b21d36e3dae2497e6bea15534c053AASv10YgRjNpvh13sJD7xgSLP0dgKAb5Sg3P4tpY7AvJAFHpEoPxNw86/NiIRqeU4VVRSwhiMWJEatoldyNhqHxrVHq8rPiEbaTfoyPb4sNRv5WnCZ8xmbAExRleQp9kKS5ZDyQRo88f0mcs50qUmpOTRbVGk/hsrRsVpkVq+EO7jMwQ5SDfnyxvXytHsZrOLEei3p060BkZqvmUTptGvoMn"

func qlogin() {
	account := module.Account{}
	account.Cookie = cookie
	client, err := q.AuthLogin(account)
	if err != nil {
		panic(err)
	}
	q.NewClient(client)
}
func TestQLogin(t *testing.T) {
	qlogin()
}
func TestQUserInfo(t *testing.T) {
	qlogin()
	info, err := q.UserInfo()
	fmt.Println(info, err)
}
func TestQShare(t *testing.T) {
	qlogin()
	url := "https://pan.quark.cn/s/41b0c70cf83c#/list/share"
	info, err := q.GetShareInfo(url, "")
	fmt.Println("TestQShare", info, err)

	//info.SToken = "8Q47RG7Wx8x3JwNuFUeRaO7sJvNpqSyopSLJ259mzKU="
	info.FileId = "0f1f1256816846dcaf50a98f60474f9e"
	list, err := q.GetSharePageFolderList(info)
	fmt.Println(list, err)
	fileList, err := q.GetSharePageFileList(info)
	fmt.Println(fileList, err)
	pageAll, err := q.GetSharePageAll(info)
	fmt.Println(pageAll, err)
}

func TestMyList(t *testing.T) {
	qlogin()
	folder, err := q.GetMyFolder("0")
	fmt.Println(folder, err)
}
