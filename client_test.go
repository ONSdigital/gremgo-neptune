package gremgo

import (
	"regexp"
	"testing"

	"github.com/gedge/gremgo-neptune/graphson"
	. "github.com/smartystreets/goconvey/convey"
)

func TestDecode(t *testing.T) {
	for _, jsTest := range res {
		Convey("Test "+jsTest.label, t, func() {
			var myStruct interface{}
			var err error
			if jsTest.parseType == "edge" {
				myStruct, err = graphson.DeserializeEdges(jsTest.js)
			} else {
				myStruct, err = graphson.DeserializeGenericValue(jsTest.js)
			}
			if err != nil {
				if jsTest.errRegex == nil {
					So(err, ShouldBeNil)
				} else {
					So(err, shouldMatch, jsTest.errRegex)
				}
			} else {
				// err == nil
				if jsTest.errRegex != nil {
					So(myStruct, ShouldResemble, "Error matching: "+jsTest.errRegex.String())
					So(err, ShouldNotBeNil)
				} else if jsTest.keyOf != nil {

				} else if jsTest.res != nil {
					So(myStruct, ShouldResemble, jsTest.res)
				}
			}
		})
	}

}

func shouldMatch(actual interface{}, expected ...interface{}) string {
	re := expected[0].(*regexp.Regexp)
	if re.MatchString(actual.(error).Error()) {
		return ""
	}
	return "No match: expected: " + expected[0].(*regexp.Regexp).String() + "\n" +
		"               got: " + actual.(error).Error()
}

type testJSON struct {
	label     string
	js        string
	errRegex  *regexp.Regexp
	res       interface{}
	hasLen    int
	parseType string
	keyOf     *string
}

var jsonParseStringErr = regexp.MustCompile("json: cannot unmarshal string")
var jsonParseObjectErr = regexp.MustCompile("json: cannot unmarshal object")
var jsonParseArrayErr = regexp.MustCompile("json: cannot unmarshal array")
var jsonParseNumberErr = regexp.MustCompile("json: cannot unmarshal number")
var emptyString = ""
var res = []testJSON{
	{
		label: "Syntax: Empty",
		js:    ``,
		res:   graphson.GenericValue{Type: "", Value: interface{}(nil)},
		// errRegex: regexp.MustCompile("eof"),
	}, {
		label: "Syntax: No JSON delimiter",
		js:    `NOPE`,
		// errRegex: jsonParseErr,
		errRegex: regexp.MustCompile(`invalid character 'N' looking for beginning of value`),
	}, {
		label:    "Syntax: Merely string",
		js:       `"NOPE"`,
		errRegex: jsonParseStringErr,
		// errRegex: regexp.MustCompile(`delim type expected, got: "string"`),
	}, {
		label:    "Syntax: Array, but not object",
		js:       `["NOPE"]`,
		errRegex: jsonParseArrayErr,
		// errRegex: regexp.MustCompile(`expected: "{", got: "\["`),
	}, {
		label: "Semantic: expect `@type` key",
		js:    `{"not@type":1234}`,
		keyOf: &emptyString,
		// errRegex: jsonParseObjectErr,
		// errRegex: regexp.MustCompile(`expected: "@type", got: "not@type"`),
	}, {
		label:    "Semantic: expect `string` value",
		js:       `{"@type":1234}`,
		errRegex: jsonParseNumberErr,
		// errRegex: regexp.MustCompile(`expected: "string", got: "json.Number"`),
	}, {
		label: "Semantic: expect limited range of values",
		js:    `{"@type":"badType"}`,
		// errRegex: jsonParseObjectErr,
		res: graphson.GenericValue{Type: "badType", Value: interface{}(nil)},
		// errRegex: regexp.MustCompile(`expected one of: \[.*\], got: "badType"`),
	}, {
		label: "Semantic: expect object not number",
		js:    `{"@type":"g:List","bar":123}`,
		res:   graphson.GenericValue{Type: "g:List", Value: interface{}(nil)},
		// errRegex: jsonParseObjectErr,
		// errRegex: regexp.MustCompile(`delim type expected, got: "json.Number"`),
	}, {
		label: "Success: minimal",
		js:    `{"@type":"g:List","@value":[{"@type":"g:Vertex","@value":{"id":"9eb43824-1a96-da70-373e-cfa846a8ef2c","label":"person","properties":{"id2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":-1442799497},"value":"81291978-cb66-4645-9124-d6248435af1c","label":"id2"}}]}}}]}`,
		res:   graphson.GenericValue{Type: "g:List", Value: []interface{}{map[string]interface{}{"@type": "g:Vertex", "@value": map[string]interface{}{"id": "9eb43824-1a96-da70-373e-cfa846a8ef2c", "label": "person", "properties": map[string]interface{}{"id2": []interface{}{map[string]interface{}{"@type": "g:VertexProperty", "@value": map[string]interface{}{"id": map[string]interface{}{"@type": "g:Int32", "@value": -1.442799497e+09}, "label": "id2", "value": "81291978-cb66-4645-9124-d6248435af1c"}}}}}}}},
	}, {
		label:  "Success: large",
		hasLen: 64,
		js:     `{"@type":"g:List","@value":[{"@type":"g:Vertex","@value":{"id":"50b2c521-8a91-457f-4a82-eb4d29c6ef67","label":"person","properties":{"name":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":2022169419},"value":"Dave","label":"name"}}]}}},{"@type":"g:Vertex","@value":{"id":"00b35fd6-b65e-4a3c-cd3b-a7cfbc93001c","label":"instance"}},{"@type":"g:Vertex","@value":{"id":"60b364f5-2dec-b3fe-6cce-54cf6b5ebf44","label":"_code_list::_instance_multiple_label_test","properties":{"name":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":-1355660727},"value":"CPIH","label":"name"}}]}}},{"@type":"g:Vertex","@value":{"id":"cab35fea-76ee-ca2f-61c7-4f3b729a2e63","label":"instance","properties":{"value":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":-700912338},"value":"v","label":"value"}}]}}},{"@type":"g:Vertex","@value":{"id":"b6b35fec-6e7b-d99d-381e-0b236cf517ef","label":"instance","properties":{"value":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":-987478415},"value":"v","label":"value"}}]}}},{"@type":"g:Vertex","@value":{"id":"82b364eb-f801-40db-8c30-51d4011e01e9","label":"person"}},{"@type":"g:Vertex","@value":{"id":"3ab364f2-d577-39d6-f6c9-b6718024117c","label":"_instance_1234","properties":{"name":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":1185796348},"value":"CPIH","label":"name"}}]}}},{"@type":"g:Vertex","@value":{"id":"f8b364ec-f83e-b568-3394-fb4c2260d4ff","label":"_instance_1234"}},{"@type":"g:Vertex","@value":{"id":"a4b41eaa-0581-510f-358e-4230a92309be","label":"person","properties":{"name":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":-874962468},"value":"Dave","label":"name"}}]}}},{"@type":"g:Vertex","@value":{"id":"f2b42d94-4e50-f528-5b16-e7c7e9a91d62","label":"person","properties":{"id2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":-945710461},"value":"4852469d-0fe3-4a20-9e38-6ea62583c564","label":"id2"}}]}}},{"@type":"g:Vertex","@value":{"id":"a8b42d9c-47e5-8e67-9399-06fd0f1f8292","label":"person","properties":{"id2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":-931853173},"value":"928b6027-33c9-4eab-ad3c-90115a480c7e","label":"id2"}}]}}},{"@type":"g:Vertex","@value":{"id":"c8b42d9d-0033-010e-f8ad-2727a6a5190b","label":"person","properties":{"id2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":480515187},"value":"2becf848-f9cd-415c-9dfb-304c862c1bc8","label":"id2"}}]}}},{"@type":"g:Vertex","@value":{"id":"beb42dbf-f0bf-fcb4-46b9-9efb65ddf8da","label":"person","properties":{"id2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":-336688873},"value":"86c81f61-fda1-4ada-8bfe-1d13b1b95040","label":"id2"}}]}}},{"@type":"g:Vertex","@value":{"id":"68b4354c-3dc7-d9c7-f0db-1c1341dc5bd4","label":"person","properties":{"id2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":-430940865},"value":"5e1699af-7992-4d2b-a886-570aa2de04d6","label":"id2"}}]}}},{"@type":"g:Vertex","@value":{"id":"48b43567-60c8-87db-e854-0cf5b5e4df41","label":"person","properties":{"id2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":-391968982},"value":"0242ec48-dcd9-42a3-b61c-76c58ade83c9","label":"id2"}}]}}},{"@type":"g:Vertex","@value":{"id":"14b437e8-4be8-ea07-471a-59be393cd2a1","label":"person","properties":{"id2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":957108541},"value":"6f20ef58-6473-447f-a8ac-f4686b1ba7ec","label":"id2"}}]}}},{"@type":"g:Vertex","@value":{"id":"d2b43813-93b4-d433-87ed-7fb49f16944a","label":"person","properties":{"id2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":1552188650},"value":"e566d36b-9fdf-473f-96ea-497b9221dbc8","label":"id2"}}]}}},{"@type":"g:Vertex","@value":{"id":"74b4381d-2364-8868-a989-6b192e4a62d8","label":"person","properties":{"id2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":-247024192},"value":"0ca78a9d-638d-44a8-9d43-404958713b74","label":"id2"}}]}}},{"@type":"g:Vertex","@value":{"id":"92b4381f-1d95-09a0-8b15-20f75f8fb97b","label":"person","properties":{"id2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":-102276262},"value":"62fa065d-fec5-4415-a312-5438677c531b","label":"id2"}}]}}},{"@type":"g:Vertex","@value":{"id":"ecb43821-78fb-66ae-5e71-45eb9ed536a9","label":"person","properties":{"id2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":-951911443},"value":"1884430f-2219-465b-85c0-f62c997dfd7f","label":"id2"}}]}}},{"@type":"g:Vertex","@value":{"id":"92b43826-1efd-9fa8-0051-304bb6568875","label":"person","properties":{"id2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":883602715},"value":"77cd85aa-c75b-4523-857b-b22f9229f262","label":"id2"}}]}}},{"@type":"g:Vertex","@value":{"id":"5ab4383c-6ace-a19c-2f8c-194ded40f3df","label":"person","properties":{"id2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":1103307885},"value":"8c3b653b-87d9-46ee-b2b7-4958968deb4a","label":"id2"}}]}}},{"@type":"g:Vertex","@value":{"id":"04b4383f-cba4-b466-966e-eb6403a58fd6","label":"person","properties":{"id2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":2011516855},"value":"480521f4-b02b-4d3b-b3a4-b5106bdfae40","label":"id2"}}]}}},{"@type":"g:Vertex","@value":{"id":"52b43846-6d64-30f7-6c72-59c6221b20a0","label":"person","properties":{"id2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":-177887341},"value":"ef9d51cb-fff9-4ce3-98e5-25a54dd04f78","label":"id2"}}]}}},{"@type":"g:Vertex","@value":{"id":"26b43856-cc5b-dc9f-d0b2-d739d82b67e9","label":"person","properties":{"id2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":1115666503},"value":"00000000-0000-0000-0000-000000000000","label":"id2"}}]}}},{"@type":"g:Vertex","@value":{"id":"38b43874-d4d4-e7af-8d4f-0d9489e02a04","label":"person","properties":{"id":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":157579180},"value":"17ae1ea0-5c4e-409d-a5bf-78e423f387e9","label":"id"}}]}}},{"@type":"g:Vertex","@value":{"id":"deb41eaa-b27b-0dda-ea20-55d95109cddd","label":"person","properties":{"name":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":-1996200004},"value":"Dave","label":"name"}}]}}},{"@type":"g:Vertex","@value":{"id":"deb42d62-950d-55ee-9da2-d997177dd504","label":"person","properties":{"id2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":1969569343},"value":"e3a0ece8-17a2-46be-b8b2-dd7e9fe50fd4","label":"id2"}}]}}},{"@type":"g:Vertex","@value":{"id":"4cb42d9a-b425-e8cb-2b87-fdbf4a7b9e93","label":"person","properties":{"id2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":-1015474976},"value":"44409e12-e07c-4dd3-bf7f-d8a07cbc015e","label":"id2"}}]}}},{"@type":"g:Vertex","@value":{"id":"0ab42d9c-924b-8828-fa38-b3c4e8702a73","label":"person","properties":{"id2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":-1853031291},"value":"1d24864a-2878-4e02-9bc8-1e3ae8d4881f","label":"id2"}}]}}},{"@type":"g:Vertex","@value":{"id":"5cb42d9e-d6c4-6b05-d37f-a94d919ab4a1","label":"person","properties":{"id2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":-1200998084},"value":"bc3fedc6-c5eb-4097-b6b5-a7882bec907d","label":"id2"}}]}}},{"@type":"g:Vertex","@value":{"id":"c8b42dc7-e473-70c2-3b34-c1a3dacbb25d","label":"person","properties":{"id2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":-1443194526},"value":"d7abf897-7a97-49c7-9e04-89c9dd3b6873","label":"id2"}}]}}},{"@type":"g:Vertex","@value":{"id":"eeb4301a-0f9b-56d4-dfd9-0c3bf5899111","label":"person","properties":{"id2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":418143122},"value":"d3d8a6c0-be5f-4a63-b0d5-16298a4900b4","label":"id2"}}]}}},{"@type":"g:Vertex","@value":{"id":"b6b437e5-c2a7-91e0-db2d-bb5ded3ab221","label":"person","properties":{"id2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":-1989625351},"value":"3585d7b7-0f72-4754-89fb-e4d6e68989cd","label":"id2"}}]}}},{"@type":"g:Vertex","@value":{"id":"98b43814-dfd6-d409-52d0-d99a57ebbce5","label":"person","properties":{"id2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":-853592014},"value":"cabd797c-4598-461a-8f99-2f77a9b8f9dc","label":"id2"}}]}}},{"@type":"g:Vertex","@value":{"id":"66b4381d-3540-246f-70ad-5d5f19793bbe","label":"person","properties":{"id2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":-236034092},"value":"fcd30bb5-4562-4c47-8e3a-cce3958d83bf","label":"id2"}}]}}},{"@type":"g:Vertex","@value":{"id":"3eb4381f-e2e9-43ff-110d-b97864a23f40","label":"person","properties":{"id2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":1545307140},"value":"1e6bb5aa-4944-4262-8d24-57d28f1b2a91","label":"id2"}}]}}},{"@type":"g:Vertex","@value":{"id":"78b43823-b85d-a479-661a-6acad25b9aa2","label":"person","properties":{"id2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":-1208335402},"value":"6718d33b-3b8c-4d69-ba51-3bfda729ea9e","label":"id2"}}]}}},{"@type":"g:Vertex","@value":{"id":"28b43856-32be-e9d5-d2d7-ef6b63e586ba","label":"person","properties":{"id2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":-1075683206},"value":"0a8b8184-0d07-4778-8d0a-13c62edf3aba","label":"id2"}}]}}},{"@type":"g:Vertex","@value":{"id":"fab43872-8eb3-58e2-b108-39a128710162","label":"person","properties":{"id":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":213618306},"value":"12569838-fd96-4db6-ae84-cb8f8858ee2f","label":"id"}}]}}},{"@type":"g:Vertex","@value":{"id":"c8b43879-2ca1-2ae0-8a72-5fcfe26f46bf","label":"person","properties":{"id":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":1036414854},"value":"2bbf75a6-2071-4907-b6ed-0af1139b8656","label":"id"}}]}}},{"@type":"g:Vertex","@value":{"id":"70b42d62-9145-345b-7001-ca9ce89d2ea8","label":"person","properties":{"id2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":-1798918806},"value":"5d0c479f-0e8a-427d-a33c-f3f77b8870a2","label":"id2"}}]}}},{"@type":"g:Vertex","@value":{"id":"9eb42d9a-b410-9638-2f2f-3e62644a5672","label":"person","properties":{"id2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":-873971625},"value":"6297c7da-6e6e-484f-ba9b-b5338cccc38d","label":"id2"}}]}}},{"@type":"g:Vertex","@value":{"id":"d2b42d9c-9237-fc9e-2f3d-d24f5a713a8b","label":"person","properties":{"id2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":526335310},"value":"91f72ca0-dc08-4cf7-852e-9c8c29a45959","label":"id2"}}]}}},{"@type":"g:Vertex","@value":{"id":"2cb42d9e-d6b0-7356-35e3-d19945c58e0e","label":"person","properties":{"id2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":-1959412460},"value":"735e0164-d6d0-4035-bf0f-ac07e571129f","label":"id2"}}]}}},{"@type":"g:Vertex","@value":{"id":"34b42dc7-e45c-c411-ec5e-984cd0f25c85","label":"person","properties":{"id2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":1224200745},"value":"721d48b8-a050-4bfa-9711-26a2a30bbb72","label":"id2"}}]}}},{"@type":"g:Vertex","@value":{"id":"ceb4301a-0f87-7107-9aef-26d2cd442777","label":"person","properties":{"id2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":-817285502},"value":"3859b9a8-f303-4a47-ba86-a85aad6a6858","label":"id2"}}]}}},{"@type":"g:Vertex","@value":{"id":"4cb437e5-bee8-fa5d-6230-bec46790a134","label":"person","properties":{"id2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":-521452030},"value":"70f02244-e46e-4799-8c6f-9f733ad8768b","label":"id2"}}]}}},{"@type":"g:Vertex","@value":{"id":"eab43813-8fdc-6a62-7684-387ae4638a96","label":"person","properties":{"id2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":1754640449},"value":"16e2c2b6-6e94-4654-9e3f-2214a3d19972","label":"id2"}}]}}},{"@type":"g:Vertex","@value":{"id":"28b4381d-2388-d636-d6c0-68c4fc5f7f29","label":"person","properties":{"id2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":-1933550086},"value":"f1f32d14-9e20-4185-b9ae-13d408ba45fb","label":"id2"}}]}}},{"@type":"g:Vertex","@value":{"id":"92b4381f-1d70-4bb0-5f15-0ea26c54479d","label":"person","properties":{"id2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":-229722235},"value":"75a5e8ea-ad97-4de5-9fd7-c7ce5308f649","label":"id2"}}]}}},{"@type":"g:Vertex","@value":{"id":"88b43821-791e-0d88-de71-7afede23b1ef","label":"person","properties":{"id2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":736913151},"value":"7ef5b170-c897-4c53-a900-a55d8e48d1be","label":"id2"}}]}}},{"@type":"g:Vertex","@value":{"id":"9cb43825-9972-bdd9-e7fb-ff558558a136","label":"person","properties":{"id2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":-1808210635},"value":"0dfd12f3-08dd-4242-bc9e-766a5b6c31da","label":"id2"}}]}}},{"@type":"g:Vertex","@value":{"id":"f8b4383c-1f91-a8cd-ae8a-5402cff32c52","label":"person","properties":{"id2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":674726809},"value":"56703150-20f6-4ab6-9195-5337497357fb","label":"id2"}}]}}},{"@type":"g:Vertex","@value":{"id":"c6b4383c-97a6-b052-3eb2-6bc0a6a1594d","label":"person","properties":{"id2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":-1687954784},"value":"9ef3b139-f371-404a-a822-a10b54a59b01","label":"id2"}}]}}},{"@type":"g:Vertex","@value":{"id":"6eb43840-f4ad-4d65-6437-9224cf70fc28","label":"person","properties":{"id2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":-1628839893},"value":"e2ccfecf-8a41-4f21-8876-81b69686acfd","label":"id2"}}]}}},{"@type":"g:Vertex","@value":{"id":"ceb43846-c843-0632-de52-f38abedd2ee4","label":"person","properties":{"id2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":-1094091875},"value":"4606e985-781a-4b35-925d-2c3f26c519d4","label":"id2"}}]}}},{"@type":"g:Vertex","@value":{"id":"50b4386d-2f0e-48a3-f5f5-4f98c439faeb","label":"person","properties":{"id":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":-1356254619},"value":"58635221-3305-4e3f-8ea8-e2268aea0d14","label":"id"}}]}}},{"@type":"g:Vertex","@value":{"id":"f6b43879-22dd-9569-b04a-b0487d92889c","label":"person","properties":{"id":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":-2074123711},"value":"4f7476ba-5cba-4480-8809-b3ee9aef5510","label":"id"}}]}}},{"@type":"g:Vertex","@value":{"id":"1cb42d94-4e66-97e3-239b-c8255d102a26","label":"person","properties":{"id2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":-1240173745},"value":"5a3ab033-0446-4c77-b8c7-68512e28be12","label":"id2"}}]}}},{"@type":"g:Vertex","@value":{"id":"6eb42d9c-4b28-fa5a-6b1b-d20e01f43117","label":"person","properties":{"id2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":1566880523},"value":"63f17e63-e542-45fb-b60b-a2de3da2a45f","label":"id2"}}]}}},{"@type":"g:Vertex","@value":{"id":"30b42d9d-0046-64fd-5850-053f42fc7e91","label":"person","properties":{"id2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":-2054863227},"value":"87c577c8-e2a7-4dd5-a62d-f0312293f812","label":"id2"}}]}}},{"@type":"g:Vertex","@value":{"id":"c8b42dbf-f0d3-77e7-3932-a00b35c813ba","label":"person","properties":{"id2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":1682216718},"value":"40eb2ca6-fb47-4f26-8985-7e4edbfbb4a0","label":"id2"}}]}}},{"@type":"g:Vertex","@value":{"id":"92b4354c-4186-acff-5f6b-0e288f26037d","label":"person","properties":{"id2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":-826943128},"value":"0196cde6-6667-40a0-8c3c-e2dd3a024e59","label":"id2"}}]}}}]}`,
	}, {
		label:  "Success: medium",
		hasLen: 9,
		js:     `{"@type":"g:List","@value":[{"@type":"g:Vertex","@value":{"id":"e0b43567-60dc-4c68-c4fe-b2ab1554893c","label":"person","properties":{"id2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":268041828},"value":"26b21a20-e1dc-42d7-a74f-97d35a5dd825","label":"id2"}}]}}},{"@type":"g:Vertex","@value":{"id":"52b437e8-4bfb-9f2d-099f-d43a7c02f908","label":"person","properties":{"id2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":798970713},"value":"0d04a43b-05f7-43bc-913a-27144174745a","label":"id2"}}]}}},{"@type":"g:Vertex","@value":{"id":"acb43814-dfb3-6b2a-ae4d-77dc6e9a8f99","label":"person","properties":{"id2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":1056643770},"value":"3b29404f-2aed-47d6-a3fc-ce36e9cfb912","label":"id2"}}]}}},{"@type":"g:Vertex","@value":{"id":"04b4381d-3562-c7d0-2180-57c1ba1fcff5","label":"person","properties":{"id2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":264990346},"value":"124ddc30-ddfa-4ac9-a4a6-a846274e5d3c","label":"id2"}}]}}},{"@type":"g:Vertex","@value":{"id":"8cb4381f-e2c8-14f4-2164-52ddf3a311d4","label":"person","properties":{"id2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":1142771785},"value":"70c8672c-6234-4f4b-b3f1-6068268fe871","label":"id2"}}]}}},{"@type":"g:Vertex","@value":{"id":"9eb43824-1a96-da70-373e-cfa846a8ef2c","label":"person","properties":{"id2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":-1442799497},"value":"81291978-cb66-4645-9124-d6248435af1c","label":"id2"}}]}}},{"@type":"g:Vertex","@value":{"id":"2eb43857-8e7f-77f5-f6ff-9d9daa9f4247","label":"person","properties":{"id":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":2021664182},"value":"388e4811-8479-4576-ac90-e3ebc6e14d97","label":"id"}}]}}},{"@type":"g:Vertex","@value":{"id":"dab43878-a010-e273-3268-45db27dbb381","label":"person","properties":{"id":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":-879701235},"value":"1e11d3dc-8a73-46dc-97b4-bc857b1144e6","label":"id"}}]}}},{"@type":"g:Vertex","@value":{"id":"e8b43879-c167-6905-35b7-df71c384a085","label":"person","properties":{"id":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":-1279362678},"value":"971de3fc-8699-4013-9cef-e396a72ea251","label":"id"}}]}}}]}`,
	}, {
		label: "Success: Edges two",
		js: `[   {"@type":"g:Edge","@value":{"id":"123","label":"label","inVLabel":"inVLabel","outVLabel":"outVLabel","inV":"1234","outV":"12345","properties":{"test":{"@type":"g:Property","@value":{"key":"test","value":{"@type":"g:Int32", "@value": 3}}}}}}, ` +
			`{"@type":"g:Edge","@value":{"id":"122","label":"label","inVLabel":"inVLabel","outVLabel":"outVLabel","inV":"1223","outV":"12234","properties":{"test":{"@type":"g:Property","@value":{"key":"test","value":{"@type":"g:Int32", "@value": 23}}}}}}]`,
		hasLen:    2,
		parseType: "edge",
	},
}

func TestGremlin(t *testing.T) {
	type testGremlin struct {
		title       string
		label       string
		input       interface{}
		expectAdd   string
		expectGet   string
		expectError error
	}

	type StructSane struct {
		Id         string
		Prop       string `graph:"prop,string"`
		IgnoreProp string
	}
	type StructSaneNoId struct {
		Prop       string `graph:"prop,string"`
		IgnoreProp string
	}
	type StructNoTags struct {
		Prop  string
		Prop2 string
	}
	type StructOtherId struct {
		Idee string `graph:"id,id"`
		Prop string `graph:"prop,string"`
	}

	type StructBool struct {
		Prop bool `graph:"prop,bool"`
	}

	res := []testGremlin{
		{
			title:     "simple",
			input:     StructSane{Id: "simple-id", Prop: "prop-val", IgnoreProp: "ignore-prop"},
			label:     "laybull",
			expectAdd: "addV('laybull').property(id,'simple-id').property('prop','prop-val')",
			expectGet: "V('laybull').hasId('simple-id').has('prop','prop-val')",
		},
		{
			title:     "escaped prop",
			input:     StructSane{Id: "simple-id", Prop: "prop-o'val", IgnoreProp: "ignore-prop"},
			label:     "escapee",
			expectAdd: `addV('escapee').property(id,'simple-id').property('prop','prop-o\'val')`,
			expectGet: `V('escapee').hasId('simple-id').has('prop','prop-o\'val')`,
		},
		{
			title:     "no-id simple",
			input:     StructSaneNoId{Prop: "prop-val", IgnoreProp: "ignore-prop"},
			label:     "no-eye-dee",
			expectAdd: "addV('no-eye-dee').property('prop','prop-val')",
			expectGet: "V('no-eye-dee').has('prop','prop-val')",
		},
		{
			title:     "no-id escaped",
			input:     StructSaneNoId{Prop: "prop-o'val", IgnoreProp: "ignore-prop"},
			label:     "no-eye-dee-esc",
			expectAdd: `addV('no-eye-dee-esc').property('prop','prop-o\'val')`,
			expectGet: `V('no-eye-dee-esc').has('prop','prop-o\'val')`,
		},
		{
			title:       "no-tags error",
			input:       StructNoTags{Prop: "p", Prop2: "p2"},
			label:       "no-tags",
			expectError: ErrorNoGraphTags,
		},
		{
			title:     "check bool type",
			input:     StructBool{Prop: true},
			label:     "typer",
			expectAdd: `addV('typer').property('prop',true)`,
			expectGet: `V('typer').has('prop',true)`,
		},
		{
			title:     "check non-Id ID",
			input:     StructOtherId{Idee: "idee-id", Prop: "prop-val"},
			label:     "other-field",
			expectAdd: `addV('other-field').property(id,'idee-id').property('prop','prop-val')`,
			expectGet: `V('other-field').hasId('idee-id').has('prop','prop-val')`,
		},
	}
	for _, gTest := range res {
		Convey("Test "+gTest.title, t, func() {
			var err error
			var outAdd, outGet string
			if outAdd, outGet, err = GremlinForVertex(gTest.label, gTest.input); err != nil {
				if gTest.expectError == nil {
					So(err, ShouldBeNil)
				}
			} else {
				// err is nil
				if gTest.expectError != nil {
					So(err, ShouldNotBeNil)
				} else {
					So(outAdd, ShouldEqual, gTest.expectAdd)
					So(outGet, ShouldEqual, gTest.expectGet)
				}
			}
		})
	}
}
