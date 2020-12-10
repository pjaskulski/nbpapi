package nbpapi

var mockTableA string = `[{"table":"A","no":"238/A/NBP/2020","effectiveDate":"2020-12-07","rates":
[{"currency":"bat (Tajlandia)","code":"THB","mid":0.1224},{"currency":"dolar amerykański","code":"USD","mid":3.7001},
{"currency":"dolar australijski","code":"AUD","mid":2.7326},{"currency":"dolar Hongkongu","code":"HKD","mid":0.4774},
{"currency":"dolar kanadyjski","code":"CAD","mid":2.8863},{"currency":"dolar nowozelandzki","code":"NZD","mid":2.5952},
{"currency":"dolar singapurski","code":"SGD","mid":2.7643},{"currency":"euro","code":"EUR","mid":4.4745},
{"currency":"forint (Węgry)","code":"HUF","mid":0.012429},{"currency":"frank szwajcarski","code":"CHF","mid":4.1417},
{"currency":"funt szterling","code":"GBP","mid":4.9089},{"currency":"hrywna (Ukraina)","code":"UAH","mid":0.1311},
{"currency":"jen (Japonia)","code":"JPY","mid":0.035478},{"currency":"korona czeska","code":"CZK","mid":0.1685},
{"currency":"korona duńska","code":"DKK","mid":0.6011},{"currency":"korona islandzka","code":"ISK","mid":0.029418},
{"currency":"korona norweska","code":"NOK","mid":0.4170},{"currency":"korona szwedzka","code":"SEK","mid":0.4357},
{"currency":"kuna (Chorwacja)","code":"HRK","mid":0.5934},{"currency":"lej rumuński","code":"RON","mid":0.9181},
{"currency":"lew (Bułgaria)","code":"BGN","mid":2.2878},{"currency":"lira turecka","code":"TRY","mid":0.4728},
{"currency":"nowy izraelski szekel","code":"ILS","mid":1.1302},{"currency":"peso chilijskie","code":"CLP","mid":0.004974},
{"currency":"peso filipińskie","code":"PHP","mid":0.0769},{"currency":"peso meksykańskie","code":"MXN","mid":0.1857},
{"currency":"rand (Republika Południowej Afryki)","code":"ZAR","mid":0.2426},{"currency":"real (Brazylia)","code":"BRL","mid":0.7176},
{"currency":"ringgit (Malezja)","code":"MYR","mid":0.9087},{"currency":"rubel rosyjski","code":"RUB","mid":0.0497},
{"currency":"rupia indonezyjska","code":"IDR","mid":0.00026232},{"currency":"rupia indyjska","code":"INR","mid":0.050078},
{"currency":"won południowokoreański","code":"KRW","mid":0.003405},{"currency":"yuan renminbi (Chiny)","code":"CNY","mid":0.5658},
{"currency":"SDR (MFW)","code":"XDR","mid":5.2966}]}]`

var mockRangeOfTablesA string = `[{"table":"A","no":"223/A/NBP/2020","effectiveDate":"2020-11-16","rates":
[{"currency":"bat (Tajlandia)","code":"THB","mid":0.1249},{"currency":"dolar amerykański","code":"USD","mid":3.7782},
{"currency":"dolar australijski","code":"AUD","mid":2.7539},{"currency":"dolar Hongkongu","code":"HKD","mid":0.4873},
{"currency":"dolar kanadyjski","code":"CAD","mid":2.8805},{"currency":"dolar nowozelandzki","code":"NZD","mid":2.5973},
{"currency":"dolar singapurski","code":"SGD","mid":2.8070},{"currency":"euro","code":"EUR","mid":4.4753},
{"currency":"forint (Węgry)","code":"HUF","mid":0.012499},{"currency":"frank szwajcarski","code":"CHF","mid":4.1463},
{"currency":"funt szterling","code":"GBP","mid":4.9761},{"currency":"hrywna (Ukraina)","code":"UAH","mid":0.1344},
{"currency":"jen (Japonia)","code":"JPY","mid":0.036139},{"currency":"korona czeska","code":"CZK","mid":0.1697},
{"currency":"korona duńska","code":"DKK","mid":0.6010},{"currency":"korona islandzka","code":"ISK","mid":0.027745},
{"currency":"korona norweska","code":"NOK","mid":0.4145},{"currency":"korona szwedzka","code":"SEK","mid":0.4362},
{"currency":"kuna (Chorwacja)","code":"HRK","mid":0.5915},{"currency":"lej rumuński","code":"RON","mid":0.9187},
{"currency":"lew (Bułgaria)","code":"BGN","mid":2.2882},{"currency":"lira turecka","code":"TRY","mid":0.4889},
{"currency":"nowy izraelski szekel","code":"ILS","mid":1.1247},{"currency":"peso chilijskie","code":"CLP","mid":0.004925},
{"currency":"peso filipińskie","code":"PHP","mid":0.0784},{"currency":"peso meksykańskie","code":"MXN","mid":0.1864},
{"currency":"rand (Republika Południowej Afryki)","code":"ZAR","mid":0.2449},{"currency":"real (Brazylia)","code":"BRL","mid":0.6923},
{"currency":"ringgit (Malezja)","code":"MYR","mid":0.9178},{"currency":"rubel rosyjski","code":"RUB","mid":0.0491},
{"currency":"rupia indonezyjska","code":"IDR","mid":0.00026777},{"currency":"rupia indyjska","code":"INR","mid":0.050709},
{"currency":"won południowokoreański","code":"KRW","mid":0.003411},{"currency":"yuan renminbi (Chiny)","code":"CNY","mid":0.5742},
{"currency":"SDR (MFW)","code":"XDR","mid":5.3780}]},
{"table":"A","no":"224/A/NBP/2020","effectiveDate":"2020-11-17","rates":
[{"currency":"bat (Tajlandia)","code":"THB","mid":0.1256},{"currency":"dolar amerykański","code":"USD","mid":3.7877},
{"currency":"dolar australijski","code":"AUD","mid":2.7758},{"currency":"dolar Hongkongu","code":"HKD","mid":0.4886},
{"currency":"dolar kanadyjski","code":"CAD","mid":2.8965},{"currency":"dolar nowozelandzki","code":"NZD","mid":2.6126},
{"currency":"dolar singapurski","code":"SGD","mid":2.8198},{"currency":"euro","code":"EUR","mid":4.4953},
{"currency":"forint (Węgry)","code":"HUF","mid":0.01244},{"currency":"frank szwajcarski","code":"CHF","mid":4.1594},
{"currency":"funt szterling","code":"GBP","mid":5.0092},{"currency":"hrywna (Ukraina)","code":"UAH","mid":0.1348},
{"currency":"jen (Japonia)","code":"JPY","mid":0.036317},{"currency":"korona czeska","code":"CZK","mid":0.1698},
{"currency":"korona duńska","code":"DKK","mid":0.6036},{"currency":"korona islandzka","code":"ISK","mid":0.027869},
{"currency":"korona norweska","code":"NOK","mid":0.4188},{"currency":"korona szwedzka","code":"SEK","mid":0.4401},
{"currency":"kuna (Chorwacja)","code":"HRK","mid":0.5941},{"currency":"lej rumuński","code":"RON","mid":0.9222},
{"currency":"lew (Bułgaria)","code":"BGN","mid":2.2984},{"currency":"lira turecka","code":"TRY","mid":0.4865},
{"currency":"nowy izraelski szekel","code":"ILS","mid":1.1295},{"currency":"peso chilijskie","code":"CLP","mid":0.004941},
{"currency":"peso filipińskie","code":"PHP","mid":0.0785},{"currency":"peso meksykańskie","code":"MXN","mid":0.1861},
{"currency":"rand (Republika Południowej Afryki)","code":"ZAR","mid":0.2456},{"currency":"real (Brazylia)","code":"BRL","mid":0.6993},
{"currency":"ringgit (Malezja)","code":"MYR","mid":0.9228},{"currency":"rubel rosyjski","code":"RUB","mid":0.0495},
{"currency":"rupia indonezyjska","code":"IDR","mid":0.00026949},{"currency":"rupia indyjska","code":"INR","mid":0.050886},
{"currency":"won południowokoreański","code":"KRW","mid":0.003427},{"currency":"yuan renminbi (Chiny)","code":"CNY","mid":0.5780},
{"currency":"SDR (MFW)","code":"XDR","mid":5.4017}]}]`

var mockTableLast5 string = `[{"table":"A","no":"234/A/NBP/2020","effectiveDate":"2020-12-01","rates":
[{"currency":"bat (Tajlandia)","code":"THB","mid":0.1235},{"currency":"dolar amerykański","code":"USD","mid":3.7367},
{"currency":"dolar australijski","code":"AUD","mid":2.7477},{"currency":"dolar Hongkongu","code":"HKD","mid":0.4820},
{"currency":"dolar kanadyjski","code":"CAD","mid":2.8801},{"currency":"dolar nowozelandzki","code":"NZD","mid":2.6281},
{"currency":"dolar singapurski","code":"SGD","mid":2.7866},{"currency":"euro","code":"EUR","mid":4.4769},
{"currency":"forint (Węgry)","code":"HUF","mid":0.012526},{"currency":"frank szwajcarski","code":"CHF","mid":4.1226},
{"currency":"funt szterling","code":"GBP","mid":4.9904},{"currency":"hrywna (Ukraina)","code":"UAH","mid":0.1308},
{"currency":"jen (Japonia)","code":"JPY","mid":0.035832},{"currency":"korona czeska","code":"CZK","mid":0.1706},
{"currency":"korona duńska","code":"DKK","mid":0.6015},{"currency":"korona islandzka","code":"ISK","mid":0.028192},
{"currency":"korona norweska","code":"NOK","mid":0.4220},{"currency":"korona szwedzka","code":"SEK","mid":0.4386},
{"currency":"kuna (Chorwacja)","code":"HRK","mid":0.5926},{"currency":"lej rumuński","code":"RON","mid":0.9188},
{"currency":"lew (Bułgaria)","code":"BGN","mid":2.2890},{"currency":"lira turecka","code":"TRY","mid":0.4762},
{"currency":"nowy izraelski szekel","code":"ILS","mid":1.1316},{"currency":"peso chilijskie","code":"CLP","mid":0.004895},
{"currency":"peso filipińskie","code":"PHP","mid":0.0777},{"currency":"peso meksykańskie","code":"MXN","mid":0.1859},
{"currency":"rand (Republika Południowej Afryki)","code":"ZAR","mid":0.2432},{"currency":"real (Brazylia)","code":"BRL","mid":0.7008},
{"currency":"ringgit (Malezja)","code":"MYR","mid":0.9163},{"currency":"rubel rosyjski","code":"RUB","mid":0.0491},
{"currency":"rupia indonezyjska","code":"IDR","mid":0.00026445},{"currency":"rupia indyjska","code":"INR","mid":0.050733},
{"currency":"won południowokoreański","code":"KRW","mid":0.003374},{"currency":"yuan renminbi (Chiny)","code":"CNY","mid":0.5687},
{"currency":"SDR (MFW)","code":"XDR","mid":5.3442}]},
{"table":"A","no":"235/A/NBP/2020","effectiveDate":"2020-12-02","rates":
[{"currency":"bat (Tajlandia)","code":"THB","mid":0.1224},{"currency":"dolar amerykański","code":"USD","mid":3.7038},
{"currency":"dolar australijski","code":"AUD","mid":2.7340},{"currency":"dolar Hongkongu","code":"HKD","mid":0.4778},
{"currency":"dolar kanadyjski","code":"CAD","mid":2.8649},{"currency":"dolar nowozelandzki","code":"NZD","mid":2.6181},
{"currency":"dolar singapurski","code":"SGD","mid":2.7629},{"currency":"euro","code":"EUR","mid":4.4642},
{"currency":"forint (Węgry)","code":"HUF","mid":0.01251},{"currency":"frank szwajcarski","code":"CHF","mid":4.1112},
{"currency":"funt szterling","code":"GBP","mid":4.9426},{"currency":"hrywna (Ukraina)","code":"UAH","mid":0.1303},
{"currency":"jen (Japonia)","code":"JPY","mid":0.035377},{"currency":"korona czeska","code":"CZK","mid":0.1695},
{"currency":"korona duńska","code":"DKK","mid":0.5996},{"currency":"korona islandzka","code":"ISK","mid":0.028362},
{"currency":"korona norweska","code":"NOK","mid":0.4194},{"currency":"korona szwedzka","code":"SEK","mid":0.4355},
{"currency":"kuna (Chorwacja)","code":"HRK","mid":0.5910},{"currency":"lej rumuński","code":"RON","mid":0.9163},
{"currency":"lew (Bułgaria)","code":"BGN","mid":2.2825},{"currency":"lira turecka","code":"TRY","mid":0.4728},
{"currency":"nowy izraelski szekel","code":"ILS","mid":1.1269},{"currency":"peso chilijskie","code":"CLP","mid":0.004873},
{"currency":"peso filipińskie","code":"PHP","mid":0.0771},{"currency":"peso meksykańskie","code":"MXN","mid":0.1844},
{"currency":"rand (Republika Południowej Afryki)","code":"ZAR","mid":0.2412},{"currency":"real (Brazylia)","code":"BRL","mid":0.7114},
{"currency":"ringgit (Malezja)","code":"MYR","mid":0.9085},{"currency":"rubel rosyjski","code":"RUB","mid":0.0490},
{"currency":"rupia indonezyjska","code":"IDR","mid":0.00026222},{"currency":"rupia indyjska","code":"INR","mid":0.050196},
{"currency":"won południowokoreański","code":"KRW","mid":0.003358},{"currency":"yuan renminbi (Chiny)","code":"CNY","mid":0.5642},
{"currency":"SDR (MFW)","code":"XDR","mid":5.3303}]},
{"table":"A","no":"236/A/NBP/2020","effectiveDate":"2020-12-03","rates":
[{"currency":"bat (Tajlandia)","code":"THB","mid":0.1225},{"currency":"dolar amerykański","code":"USD","mid":3.6981},
{"currency":"dolar australijski","code":"AUD","mid":2.7440},{"currency":"dolar Hongkongu","code":"HKD","mid":0.4771},
{"currency":"dolar kanadyjski","code":"CAD","mid":2.8593},{"currency":"dolar nowozelandzki","code":"NZD","mid":2.6108},
{"currency":"dolar singapurski","code":"SGD","mid":2.7648},{"currency":"euro","code":"EUR","mid":4.4789},
{"currency":"forint (Węgry)","code":"HUF","mid":0.012477},{"currency":"frank szwajcarski","code":"CHF","mid":4.1362},
{"currency":"funt szterling","code":"GBP","mid":4.9562},{"currency":"hrywna (Ukraina)","code":"UAH","mid":0.1306},
{"currency":"jen (Japonia)","code":"JPY","mid":0.035464},{"currency":"korona czeska","code":"CZK","mid":0.1692},
{"currency":"korona duńska","code":"DKK","mid":0.6017},{"currency":"korona islandzka","code":"ISK","mid":0.029065},
{"currency":"korona norweska","code":"NOK","mid":0.4187},{"currency":"korona szwedzka","code":"SEK","mid":0.4354},
{"currency":"kuna (Chorwacja)","code":"HRK","mid":0.5934},{"currency":"lej rumuński","code":"RON","mid":0.9190},
{"currency":"lew (Bułgaria)","code":"BGN","mid":2.2900},{"currency":"lira turecka","code":"TRY","mid":0.4720},
{"currency":"nowy izraelski szekel","code":"ILS","mid":1.1275},{"currency":"peso chilijskie","code":"CLP","mid":0.004892},
{"currency":"peso filipińskie","code":"PHP","mid":0.0770},{"currency":"peso meksykańskie","code":"MXN","mid":0.1851},
{"currency":"rand (Republika Południowej Afryki)","code":"ZAR","mid":0.2414},{"currency":"real (Brazylia)","code":"BRL","mid":0.7087},
{"currency":"ringgit (Malezja)","code":"MYR","mid":0.9077},{"currency":"rubel rosyjski","code":"RUB","mid":0.0493},
{"currency":"rupia indonezyjska","code":"IDR","mid":0.00026153},{"currency":"rupia indyjska","code":"INR","mid":0.050016},
{"currency":"won południowokoreański","code":"KRW","mid":0.003384},{"currency":"yuan renminbi (Chiny)","code":"CNY","mid":0.5639},
{"currency":"SDR (MFW)","code":"XDR","mid":5.3157}]},
{"table":"A","no":"237/A/NBP/2020","effectiveDate":"2020-12-04","rates":
[{"currency":"bat (Tajlandia)","code":"THB","mid":0.1219},{"currency":"dolar amerykański","code":"USD","mid":3.6765},
{"currency":"dolar australijski","code":"AUD","mid":2.7305},{"currency":"dolar Hongkongu","code":"HKD","mid":0.4743},
{"currency":"dolar kanadyjski","code":"CAD","mid":2.8596},{"currency":"dolar nowozelandzki","code":"NZD","mid":2.5928},
{"currency":"dolar singapurski","code":"SGD","mid":2.7600},{"currency":"euro","code":"EUR","mid":4.4732},
{"currency":"forint (Węgry)","code":"HUF","mid":0.012474},{"currency":"frank szwajcarski","code":"CHF","mid":4.1241},
{"currency":"funt szterling","code":"GBP","mid":4.9565},{"currency":"hrywna (Ukraina)","code":"UAH","mid":0.1300},
{"currency":"jen (Japonia)","code":"JPY","mid":0.035367},{"currency":"korona czeska","code":"CZK","mid":0.1689},
{"currency":"korona duńska","code":"DKK","mid":0.6009},{"currency":"korona islandzka","code":"ISK","mid":0.029294},
{"currency":"korona norweska","code":"NOK","mid":0.4198},{"currency":"korona szwedzka","code":"SEK","mid":0.4356},
{"currency":"kuna (Chorwacja)","code":"HRK","mid":0.5928},{"currency":"lej rumuński","code":"RON","mid":0.9179},
{"currency":"lew (Bułgaria)","code":"BGN","mid":2.2871},{"currency":"lira turecka","code":"TRY","mid":0.4734},
{"currency":"nowy izraelski szekel","code":"ILS","mid":1.1257},{"currency":"peso chilijskie","code":"CLP","mid":0.004898},
{"currency":"peso filipińskie","code":"PHP","mid":0.0765},{"currency":"peso meksykańskie","code":"MXN","mid":0.1851},
{"currency":"rand (Republika Południowej Afryki)","code":"ZAR","mid":0.2420},{"currency":"real (Brazylia)","code":"BRL","mid":0.7136},
{"currency":"ringgit (Malezja)","code":"MYR","mid":0.9055},{"currency":"rubel rosyjski","code":"RUB","mid":0.0495},
{"currency":"rupia indonezyjska","code":"IDR","mid":0.00026065},{"currency":"rupia indyjska","code":"INR","mid":0.049835},
{"currency":"won południowokoreański","code":"KRW","mid":0.003392},{"currency":"yuan renminbi (Chiny)","code":"CNY","mid":0.5630},
{"currency":"SDR (MFW)","code":"XDR","mid":5.2891}]},
{"table":"A","no":"238/A/NBP/2020","effectiveDate":"2020-12-07","rates":
[{"currency":"bat (Tajlandia)","code":"THB","mid":0.1224},{"currency":"dolar amerykański","code":"USD","mid":3.7001},
{"currency":"dolar australijski","code":"AUD","mid":2.7326},{"currency":"dolar Hongkongu","code":"HKD","mid":0.4774},
{"currency":"dolar kanadyjski","code":"CAD","mid":2.8863},{"currency":"dolar nowozelandzki","code":"NZD","mid":2.5952},
{"currency":"dolar singapurski","code":"SGD","mid":2.7643},{"currency":"euro","code":"EUR","mid":4.4745},
{"currency":"forint (Węgry)","code":"HUF","mid":0.012429},{"currency":"frank szwajcarski","code":"CHF","mid":4.1417},
{"currency":"funt szterling","code":"GBP","mid":4.9089},{"currency":"hrywna (Ukraina)","code":"UAH","mid":0.1311},
{"currency":"jen (Japonia)","code":"JPY","mid":0.035478},{"currency":"korona czeska","code":"CZK","mid":0.1685},
{"currency":"korona duńska","code":"DKK","mid":0.6011},{"currency":"korona islandzka","code":"ISK","mid":0.029418},
{"currency":"korona norweska","code":"NOK","mid":0.4170},{"currency":"korona szwedzka","code":"SEK","mid":0.4357},
{"currency":"kuna (Chorwacja)","code":"HRK","mid":0.5934},{"currency":"lej rumuński","code":"RON","mid":0.9181},
{"currency":"lew (Bułgaria)","code":"BGN","mid":2.2878},{"currency":"lira turecka","code":"TRY","mid":0.4728},
{"currency":"nowy izraelski szekel","code":"ILS","mid":1.1302},{"currency":"peso chilijskie","code":"CLP","mid":0.004974},
{"currency":"peso filipińskie","code":"PHP","mid":0.0769},{"currency":"peso meksykańskie","code":"MXN","mid":0.1857},
{"currency":"rand (Republika Południowej Afryki)","code":"ZAR","mid":0.2426},{"currency":"real (Brazylia)","code":"BRL","mid":0.7176},
{"currency":"ringgit (Malezja)","code":"MYR","mid":0.9087},{"currency":"rubel rosyjski","code":"RUB","mid":0.0497},
{"currency":"rupia indonezyjska","code":"IDR","mid":0.00026232},{"currency":"rupia indyjska","code":"INR","mid":0.050078},
{"currency":"won południowokoreański","code":"KRW","mid":0.003405},{"currency":"yuan renminbi (Chiny)","code":"CNY","mid":0.5658},
{"currency":"SDR (MFW)","code":"XDR","mid":5.2966}]}]`

var mockTableC string = `[{"table":"C","no":"238/C/NBP/2020","tradingDate":"2020-12-04","effectiveDate":"2020-12-07","rates":
[{"currency":"dolar amerykański","code":"USD","bid":3.6408,"ask":3.7144},{"currency":"dolar australijski","code":"AUD","bid":2.7081,"ask":2.7629},
{"currency":"dolar kanadyjski","code":"CAD","bid":2.8419,"ask":2.8993},{"currency":"euro","code":"EUR","bid":4.4264,"ask":4.5158},
{"currency":"forint (Węgry)","code":"HUF","bid":0.012349,"ask":0.012599},{"currency":"frank szwajcarski","code":"CHF","bid":4.0937,"ask":4.1765},
{"currency":"funt szterling","code":"GBP","bid":4.9226,"ask":5.0220},{"currency":"jen (Japonia)","code":"JPY","bid":0.035006,"ask":0.035714},
{"currency":"korona czeska","code":"CZK","bid":0.1670,"ask":0.1704},{"currency":"korona duńska","code":"DKK","bid":0.5947,"ask":0.6067},
{"currency":"korona norweska","code":"NOK","bid":0.4154,"ask":0.4238},{"currency":"korona szwedzka","code":"SEK","bid":0.4316,"ask":0.4404},
{"currency":"SDR (MFW)","code":"XDR","bid":5.2337,"ask":5.3395}]}]`

var mockTableCLast5 string = `[{"table":"C","no":"234/C/NBP/2020","tradingDate":"2020-11-30","effectiveDate":"2020-12-01","rates":
[{"currency":"dolar amerykański","code":"USD","bid":3.6925,"ask":3.7671},{"currency":"dolar australijski","code":"AUD","bid":2.7241,"ask":2.7791},
{"currency":"dolar kanadyjski","code":"CAD","bid":2.8547,"ask":2.9123},{"currency":"euro","code":"EUR","bid":4.4292,"ask":4.5186},
{"currency":"forint (Węgry)","code":"HUF","bid":0.012331,"ask":0.012581},{"currency":"frank szwajcarski","code":"CHF","bid":4.0842,"ask":4.1668},
{"currency":"funt szterling","code":"GBP","bid":4.9377,"ask":5.0375},{"currency":"jen (Japonia)","code":"JPY","bid":0.035419,"ask":0.036135},
{"currency":"korona czeska","code":"CZK","bid":0.1689,"ask":0.1723},{"currency":"korona duńska","code":"DKK","bid":0.5952,"ask":0.6072},
{"currency":"korona norweska","code":"NOK","bid":0.4194,"ask":0.4278},{"currency":"korona szwedzka","code":"SEK","bid":0.4352,"ask":0.4440},
{"currency":"SDR (MFW)","code":"XDR","bid":5.3093,"ask":5.4165}]},
{"table":"C","no":"235/C/NBP/2020","tradingDate":"2020-12-01","effectiveDate":"2020-12-02","rates":
[{"currency":"dolar amerykański","code":"USD","bid":3.6809,"ask":3.7553},{"currency":"dolar australijski","code":"AUD","bid":2.7070,"ask":2.7616},
{"currency":"dolar kanadyjski","code":"CAD","bid":2.8394,"ask":2.8968},{"currency":"euro","code":"EUR","bid":4.4151,"ask":4.5043},
{"currency":"forint (Węgry)","code":"HUF","bid":0.012409,"ask":0.012659},{"currency":"frank szwajcarski","code":"CHF","bid":4.0724,"ask":4.1546},
{"currency":"funt szterling","code":"GBP","bid":4.9144,"ask":5.0136},{"currency":"jen (Japonia)","code":"JPY","bid":0.035226,"ask":0.035938},
{"currency":"korona czeska","code":"CZK","bid":0.1680,"ask":0.1714},{"currency":"korona duńska","code":"DKK","bid":0.5931,"ask":0.6051},
{"currency":"korona norweska","code":"NOK","bid":0.4159,"ask":0.4243},{"currency":"korona szwedzka","code":"SEK","bid":0.4314,"ask":0.4402},
{"currency":"SDR (MFW)","code":"XDR","bid":5.2705,"ask":5.3769}]},
{"table":"C","no":"236/C/NBP/2020","tradingDate":"2020-12-02","effectiveDate":"2020-12-03","rates":
[{"currency":"dolar amerykański","code":"USD","bid":3.6745,"ask":3.7487},{"currency":"dolar australijski","code":"AUD","bid":2.7064,"ask":2.7610},
{"currency":"dolar kanadyjski","code":"CAD","bid":2.8410,"ask":2.8984},{"currency":"euro","code":"EUR","bid":4.4399,"ask":4.5295},
{"currency":"forint (Węgry)","code":"HUF","bid":0.012376,"ask":0.012626},{"currency":"frank szwajcarski","code":"CHF","bid":4.1025,"ask":4.1853},
{"currency":"funt szterling","code":"GBP","bid":4.8897,"ask":4.9885},{"currency":"jen (Japonia)","code":"JPY","bid":0.035162,"ask":0.035872},
{"currency":"korona czeska","code":"CZK","bid":0.1683,"ask":0.1717},{"currency":"korona duńska","code":"DKK","bid":0.5965,"ask":0.6085},
{"currency":"korona norweska","code":"NOK","bid":0.4152,"ask":0.4236},{"currency":"korona szwedzka","code":"SEK","bid":0.4315,"ask":0.4403},
{"currency":"SDR (MFW)","code":"XDR","bid":5.3013,"ask":5.4083}]},
{"table":"C","no":"237/C/NBP/2020","tradingDate":"2020-12-03","effectiveDate":"2020-12-04","rates":
[{"currency":"dolar amerykański","code":"USD","bid":3.6421,"ask":3.7157},{"currency":"dolar australijski","code":"AUD","bid":2.7098,"ask":2.7646},
{"currency":"dolar kanadyjski","code":"CAD","bid":2.8238,"ask":2.8808},{"currency":"euro","code":"EUR","bid":4.4287,"ask":4.5181},
{"currency":"forint (Węgry)","code":"HUF","bid":0.012385,"ask":0.012635},{"currency":"frank szwajcarski","code":"CHF","bid":4.0879,"ask":4.1705},
{"currency":"funt szterling","code":"GBP","bid":4.9011,"ask":5.0001},{"currency":"jen (Japonia)","code":"JPY","bid":0.035082,"ask":0.03579},
{"currency":"korona czeska","code":"CZK","bid":0.1676,"ask":0.1710},{"currency":"korona duńska","code":"DKK","bid":0.5949,"ask":0.6069},
{"currency":"korona norweska","code":"NOK","bid":0.4155,"ask":0.4239},{"currency":"korona szwedzka","code":"SEK","bid":0.4311,"ask":0.4399},
{"currency":"SDR (MFW)","code":"XDR","bid":5.2561,"ask":5.3623}]},
{"table":"C","no":"238/C/NBP/2020","tradingDate":"2020-12-04","effectiveDate":"2020-12-07","rates":
[{"currency":"dolar amerykański","code":"USD","bid":3.6408,"ask":3.7144},{"currency":"dolar australijski","code":"AUD","bid":2.7081,"ask":2.7629},
{"currency":"dolar kanadyjski","code":"CAD","bid":2.8419,"ask":2.8993},{"currency":"euro","code":"EUR","bid":4.4264,"ask":4.5158},
{"currency":"forint (Węgry)","code":"HUF","bid":0.012349,"ask":0.012599},{"currency":"frank szwajcarski","code":"CHF","bid":4.0937,"ask":4.1765},
{"currency":"funt szterling","code":"GBP","bid":4.9226,"ask":5.0220},{"currency":"jen (Japonia)","code":"JPY","bid":0.035006,"ask":0.035714},
{"currency":"korona czeska","code":"CZK","bid":0.1670,"ask":0.1704},{"currency":"korona duńska","code":"DKK","bid":0.5947,"ask":0.6067},
{"currency":"korona norweska","code":"NOK","bid":0.4154,"ask":0.4238},{"currency":"korona szwedzka","code":"SEK","bid":0.4316,"ask":0.4404},
{"currency":"SDR (MFW)","code":"XDR","bid":5.2337,"ask":5.3395}]}]`

var mockTableCXML string = `<ArrayOfExchangeRatesTable xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
<ExchangeRatesTable>
<Table>C</Table>
<No>238/C/NBP/2020</No>
<TradingDate>2020-12-04</TradingDate>
<EffectiveDate>2020-12-07</EffectiveDate>
<Rates>
<Rate>
<Currency>dolar amerykański</Currency>
<Code>USD</Code>
<Bid>3.6408</Bid>
<Ask>3.7144</Ask>
</Rate>
<Rate>
<Currency>dolar australijski</Currency>
<Code>AUD</Code>
<Bid>2.7081</Bid>
<Ask>2.7629</Ask>
</Rate>
<Rate>
<Currency>dolar kanadyjski</Currency>
<Code>CAD</Code>
<Bid>2.8419</Bid>
<Ask>2.8993</Ask>
</Rate>
<Rate>
<Currency>euro</Currency>
<Code>EUR</Code>
<Bid>4.4264</Bid>
<Ask>4.5158</Ask>
</Rate>
<Rate>
<Currency>forint (Węgry)</Currency>
<Code>HUF</Code>
<Bid>0.012349</Bid>
<Ask>0.012599</Ask>
</Rate>
<Rate>
<Currency>frank szwajcarski</Currency>
<Code>CHF</Code>
<Bid>4.0937</Bid>
<Ask>4.1765</Ask>
</Rate>
<Rate>
<Currency>funt szterling</Currency>
<Code>GBP</Code>
<Bid>4.9226</Bid>
<Ask>5.0220</Ask>
</Rate>
<Rate>
<Currency>jen (Japonia)</Currency>
<Code>JPY</Code>
<Bid>0.035006</Bid>
<Ask>0.035714</Ask>
</Rate>
<Rate>
<Currency>korona czeska</Currency>
<Code>CZK</Code>
<Bid>0.1670</Bid>
<Ask>0.1704</Ask>
</Rate>
<Rate>
<Currency>korona duńska</Currency>
<Code>DKK</Code>
<Bid>0.5947</Bid>
<Ask>0.6067</Ask>
</Rate>
<Rate>
<Currency>korona norweska</Currency>
<Code>NOK</Code>
<Bid>0.4154</Bid>
<Ask>0.4238</Ask>
</Rate>
<Rate>
<Currency>korona szwedzka</Currency>
<Code>SEK</Code>
<Bid>0.4316</Bid>
<Ask>0.4404</Ask>
</Rate>
<Rate>
<Currency>SDR (MFW)</Currency>
<Code>XDR</Code>
<Bid>5.2337</Bid>
<Ask>5.3395</Ask>
</Rate>
</Rates>
</ExchangeRatesTable>
</ArrayOfExchangeRatesTable>`
