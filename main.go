package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	const myurl = "http://ecom-multistore-core.com/index.php?dispatch=payment_notification.confirm&payment=mobilpay_ro&order_id=13388"
	const xmlbody = `<?xml version="1.0" encoding="utf-8"?> <order type="card" id="13388" timestamp="20210606124944"><signature>YHBU-E8JA-LGXR-LECV-5TSL</signature><invoice currency="RON" amount="80.00"><contact_info><billing type="person"><address><![CDATA[RO%2C+%2C+BV%2C+Bucurest%2C+Vale+Dos+Vales]]></address><email><![CDATA[stgirina%40yopmail.com]]></email></billing><shipping type="person"><address><![CDATA[RO%2C+%2C+BV%2C+Bucurest%2C+Vale+Dos+Vales]]></address><email><![CDATA[stgirina%40yopmail.com]]></email><mobile_phone><![CDATA[0712345678]]></mobile_phone></shipping></contact_info></invoice><url><return>https://staging.iqos.ro/index.php?dispatch=payment_notification.return&amp;payment=mobilpay_ro&amp;order_id=13388</return><confirm>https://staging.iqos.ro/index.php?dispatch=payment_notification.confirm&amp;payment=mobilpay_ro&amp;order_id=13388</confirm></url><mobilpay timestamp="20210606125130" crc="35f8fc625aa37f8fe4ebe4013364af40"><action>confirmed</action><customer type="person"><first_name><![CDATA[test+test]]></first_name><address><![CDATA[RO%2C+%2C+BV%2C+Bucurest%2C+Vale+Dos+Vales]]></address><email><![CDATA[stgirina%40yopmail.com]]></email><mobile_phone><![CDATA[0746010114]]></mobile_phone></customer><purchase>1233802</purchase><original_amount>80.00</original_amount><processed_amount>80.00</processed_amount><current_payment_count>1</current_payment_count><pan_masked>9****5098</pan_masked><rrn>9991622</rrn><payment_instrument_id>41679</payment_instrument_id><installments>1</installments><error code="0"><![CDATA[Tranzactia aprobata]]></error></mobilpay></order>`

	resp, err := http.Post(myurl, "text/xml", strings.NewReader(xmlbody))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println(resp.Body)
}