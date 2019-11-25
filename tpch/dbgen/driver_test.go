package dbgen

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

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

var gotLinesBuf bytes.Buffer
var expectLines = `1|155190|7706|1|17|21168.23|0.04|0.02|N|O|1996-03-13|1996-02-12|1996-03-22|DELIVER IN PERSON|TRUCK|egular courts above the|
1|67310|7311|2|36|45983.16|0.09|0.06|N|O|1996-04-12|1996-02-28|1996-04-20|TAKE BACK RETURN|MAIL|ly final dependencies: slyly bold |
1|63700|3701|3|8|13309.60|0.10|0.02|N|O|1996-01-29|1996-03-05|1996-01-31|TAKE BACK RETURN|REG AIR|riously. regular, express dep|
1|2132|4633|4|28|28955.64|0.09|0.06|N|O|1996-04-21|1996-03-30|1996-05-16|NONE|AIR|lites. fluffily even de|
1|24027|1534|5|24|22824.48|0.10|0.04|N|O|1996-03-30|1996-03-14|1996-04-01|NONE|FOB| pending foxes. slyly re|
1|15635|638|6|32|49620.16|0.07|0.02|N|O|1996-01-30|1996-02-07|1996-02-03|DELIVER IN PERSON|MAIL|arefully slyly ex|
2|106170|1191|1|38|44694.46|0.00|0.05|N|O|1997-01-28|1997-01-14|1997-02-02|TAKE BACK RETURN|RAIL|ven requests. deposits breach a|
3|4297|1798|1|45|54058.05|0.06|0.00|R|F|1994-02-02|1994-01-04|1994-02-23|NONE|AIR|ongside of the furiously brave acco|
3|19036|6540|2|49|46796.47|0.10|0.00|R|F|1993-11-09|1993-12-20|1993-11-24|TAKE BACK RETURN|RAIL| unusual accounts. eve|
3|128449|3474|3|27|39890.88|0.06|0.07|A|F|1994-01-16|1993-11-22|1994-01-23|DELIVER IN PERSON|SHIP|nal foxes wake. |
3|29380|1883|4|2|2618.76|0.01|0.06|A|F|1993-12-04|1994-01-07|1994-01-01|NONE|TRUCK|y. fluffily pending d|
3|183095|650|5|28|32986.52|0.04|0.00|R|F|1993-12-14|1994-01-10|1994-01-01|TAKE BACK RETURN|FOB|ages nag slyly pending|
3|62143|9662|6|26|28733.64|0.10|0.02|A|F|1993-10-29|1993-12-18|1993-11-04|TAKE BACK RETURN|RAIL|ges sleep after the caref|
4|88035|5560|1|30|30690.90|0.03|0.08|N|O|1996-01-10|1995-12-14|1996-01-18|DELIVER IN PERSON|REG AIR|- quickly regular packages sleep. idly|
5|108570|8571|1|15|23678.55|0.02|0.04|R|F|1994-10-31|1994-08-31|1994-11-20|NONE|AIR|ts wake furiously |
5|123927|3928|2|26|50723.92|0.07|0.08|R|F|1994-10-16|1994-09-25|1994-10-19|NONE|FOB|sts use slyly quickly special instruc|
5|37531|35|3|50|73426.50|0.08|0.03|A|F|1994-08-08|1994-10-13|1994-08-26|DELIVER IN PERSON|AIR|eodolites. fluffily unusual|
6|139636|2150|1|37|61998.31|0.08|0.03|A|F|1992-04-27|1992-05-15|1992-05-02|TAKE BACK RETURN|TRUCK|p furiously special foxes|
7|182052|9607|1|12|13608.60|0.07|0.03|N|O|1996-05-07|1996-03-13|1996-06-03|TAKE BACK RETURN|FOB|ss pinto beans wake against th|
7|145243|7758|2|9|11594.16|0.08|0.08|N|O|1996-02-01|1996-03-02|1996-02-19|TAKE BACK RETURN|SHIP|es. instructions|
7|94780|9799|3|46|81639.88|0.10|0.07|N|O|1996-01-15|1996-03-27|1996-02-03|COLLECT COD|MAIL| unusual reques|
7|163073|3074|4|28|31809.96|0.03|0.04|N|O|1996-03-21|1996-04-08|1996-04-20|NONE|FOB|. slyly special requests haggl|
7|151894|9440|5|38|73943.82|0.08|0.01|N|O|1996-02-11|1996-02-24|1996-02-18|DELIVER IN PERSON|TRUCK|ns haggle carefully ironic deposits. bl|
7|79251|1759|6|35|43058.75|0.06|0.03|N|O|1996-01-16|1996-02-23|1996-01-22|TAKE BACK RETURN|FOB|jole. excuses wake carefully alongside of |
7|157238|2269|7|5|6476.15|0.04|0.02|N|O|1996-02-10|1996-03-26|1996-02-13|NONE|FOB|ithely regula|
32|82704|7721|1|28|47227.60|0.05|0.08|N|O|1995-10-23|1995-08-27|1995-10-26|TAKE BACK RETURN|TRUCK|sleep quickly. req|
32|197921|441|2|32|64605.44|0.02|0.00|N|O|1995-08-14|1995-10-07|1995-08-27|COLLECT COD|AIR|lithely regular deposits. fluffily |
32|44161|6666|3|2|2210.32|0.09|0.02|N|O|1995-08-07|1995-10-07|1995-08-23|DELIVER IN PERSON|AIR| express accounts wake according to the|
32|2743|7744|4|4|6582.96|0.09|0.03|N|O|1995-08-04|1995-10-01|1995-09-03|NONE|REG AIR|e slyly final pac|
32|85811|8320|5|44|79059.64|0.05|0.06|N|O|1995-08-28|1995-08-20|1995-09-14|DELIVER IN PERSON|AIR|symptotes nag according to the ironic depo|
32|11615|4117|6|6|9159.66|0.04|0.03|N|O|1995-07-21|1995-09-23|1995-07-25|COLLECT COD|RAIL| gifts cajole carefully.|
33|61336|8855|1|31|40217.23|0.09|0.04|A|F|1993-10-29|1993-12-19|1993-11-08|COLLECT COD|TRUCK|ng to the furiously ironic package|
33|60519|5532|2|32|47344.32|0.02|0.05|A|F|1993-12-09|1994-01-04|1993-12-28|COLLECT COD|MAIL|gular theodolites|
33|137469|9983|3|5|7532.30|0.05|0.03|A|F|1993-12-09|1993-12-25|1993-12-23|TAKE BACK RETURN|AIR|. stealthily bold exc|
33|33918|3919|4|41|75928.31|0.09|0.00|R|F|1993-11-09|1994-01-24|1993-11-11|TAKE BACK RETURN|MAIL|unusual packages doubt caref|
34|88362|871|1|13|17554.68|0.00|0.07|N|O|1998-10-23|1998-09-14|1998-11-06|NONE|REG AIR|nic accounts. deposits are alon|
34|89414|1923|2|22|30875.02|0.08|0.06|N|O|1998-10-09|1998-10-16|1998-10-12|NONE|FOB|thely slyly p|
34|169544|4577|3|6|9681.24|0.02|0.06|N|O|1998-10-30|1998-09-20|1998-11-05|NONE|FOB|ar foxes sleep |
`

func TestMain(m *testing.M) {
	initDriver(1)

	testOrderLoader := func(order interface{}) error {
		o := order.(*Order)
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

	testLineLoader := func(order interface{}) error {
		o := order.(*Order)
		for _, line := range o.lines {
			if _, err := gotLinesBuf.WriteString(
				fmt.Sprintf("%d|%d|%d|%d|%d|%d.%02d|%d.%02d|%d.%02d|%c|%c|%s|%s|%s|%s|%s|%s|\n",
					line.oKey,
					line.partKey,
					line.suppKey,
					line.lCnt,
					line.quantity,
					line.ePrice/100, line.ePrice%100,
					line.discount/100, line.discount%100,
					line.tax/100, line.tax%100,
					line.rFlag,
					line.lStatus,
					line.sDate,
					line.cDate,
					line.rDate,
					line.shipInstruct,
					line.shipMode,
					line.comment,
				)); err != nil {
				return err
			}
		}

		return nil
	}

	*orderLoader = testOrderLoader
	*lineItemLoader = testLineLoader
	os.Exit(m.Run())
}

func TestGenOrder(t *testing.T) {
	if err := genTable(ORDER, 1, 10); err != nil {
		t.Error(err)
	}

	gotOrders := string(gotOrdersBuf.Bytes())
	if gotOrders != expectOrders {
		t.Errorf("expect:\n%s\ngot:\n%s", expectOrders, gotOrders)
	}
}

func TestGenLine(t *testing.T) {
	if err := genTable(LINE, 1, 10); err != nil {
		t.Error(err)
	}

	gotLines := string(gotLinesBuf.Bytes())
	if gotLines != expectLines {
		t.Errorf("expect:\n%s\ngot:\n%s", expectLines, gotLines)
	}
}

func TestGenOrderLine(t *testing.T) {
	if err := genTable(ORDER_LINE, 1, 10); err != nil {
		t.Error(err)
	}
	gotOrders := string(gotOrdersBuf.Bytes())
	if gotOrders != expectOrders {
		t.Errorf("expect:\n%s\ngot:\n%s", expectOrders, gotOrders)
	}
	gotLines := string(gotLinesBuf.Bytes())
	if gotLines != expectLines {
		t.Errorf("expect:\n%s\ngot:\n%s", expectLines, gotLines)
	}
}
