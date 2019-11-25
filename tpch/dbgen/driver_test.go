package dbgen

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	initDriver(1)
	os.Exit(m.Run())
}

var gotOrdersBuf bytes.Buffer
var expectOrders = `1|36901|O|173665.47|1996-01-02|5-LOW|Clerk#000000951|0|nstructions sleep furiously among |
2|78002|O|46929.18|1996-12-01|1-URGENT|Clerk#000000880|0| foxes. pending accounts at the pending, silent asymptot|
3|123314|F|193846.25|1993-10-14|5-LOW|Clerk#000000955|0|sly final accounts boost. carefully regular ideas cajole carefully. depos|
4|136777|O|32151.78|1995-10-11|5-LOW|Clerk#000000124|0|sits. slyly regular warthogs cajole. regular, regular theodolites acro|
5|44485|F|144659.20|1994-07-30|5-LOW|Clerk#000000925|0|quickly. bold deposits sleep slyly. packages use slyly|
6|55624|F|58749.59|1992-02-21|4-NOT SPECIFIED|Clerk#000000058|0|ggle. special, final requests are against the furiously specia|
7|39136|O|252004.18|1996-01-10|2-HIGH|Clerk#000000470|0|ly special requests |
32|130057|O|208660.75|1995-07-16|2-HIGH|Clerk#000000616|0|ise blithely bold, regular requests. quickly unusual dep|
33|66958|F|163243.98|1993-10-27|3-MEDIUM|Clerk#000000409|0|uriously. furiously final request|
34|61001|O|58949.67|1998-07-21|3-MEDIUM|Clerk#000000223|0|ly final packages. fluffily final deposits wake blithely ideas. spe|
`

func TestGenTable(t *testing.T) {
	orderLoader = func(o *Order) error {
		gotOrdersBuf.WriteString(fmt.Sprintf("%d|%d|%c|%d.%02d|%s|%s|%s|%d|%s|\n",
			o.oKey,
			o.custKey,
			o.status,
			o.totalPrice/100, o.totalPrice%100,
			o.date,
			o.orderPriority,
			o.clerk,
			o.shipPriority,
			o.comment))

		return nil
	}

	if err := genTable(ORDER, 1, 10); err != nil {
		t.Error(err)
	}

	gotOrders := string(gotOrdersBuf.Bytes())
	if gotOrders != expectOrders {
		t.Errorf("expect:\n%s\ngot:\n%s", expectOrders, gotOrders)
	}
}
