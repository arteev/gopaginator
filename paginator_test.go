package gopaginator

import (
	"testing"
	"fmt"
)


const templateUrl = "?page=%s"

func TestFormatUrl(t *testing.T) {
	var s,must string

	s = formatUrl("1",1,2,templateUrl);
	must = fmt.Sprintf(templateUrl,"1")
	if s != must {
		t.Errorf("Expected %q, got %q",must, s)
	}


	s = formatUrl("2",1,1,templateUrl);
	must = fmt.Sprintf(templateUrl,"1")
	if s != must {
		t.Errorf("Expected %q, got %q",must, s)
	}

	s = formatUrl("2",1,0,templateUrl);
	must = ""
	if s != must {
		t.Errorf("Expected %q, got %q",must, s)
	}

	s = formatUrl("11",1,10,templateUrl);
	must = fmt.Sprintf(templateUrl,"10")
	if s != must {
		t.Errorf("Expected %q, got %q",must, s)
	}


	s = formatUrl("<<",1,10,templateUrl);
	must = fmt.Sprintf(templateUrl,"1")
	if s != must {
		t.Errorf("Expected %q, got %q",must, s)
	}

	s = formatUrl("<<",1,-1,templateUrl);
	must =""
	if s != must {
		t.Errorf("Expected %q, got %q",must, s)
	}

	/// <
	s = formatUrl("<",1,10,templateUrl);
	must = fmt.Sprintf(templateUrl,"1")
	if s != must {
		t.Errorf("Expected %q, got %q",must, s)
	}

	s = formatUrl("<",9,10,templateUrl);
	must = fmt.Sprintf(templateUrl,"8")
	if s != must {
		t.Errorf("Expected %q, got %q",must, s)
	}


	s = formatUrl("<",0,10,templateUrl);
	must = ""
	if s != must {
		t.Errorf("Expected %q, got %q",must, s)
	}


	///  >>
	s = formatUrl(">>",1,10,templateUrl);
	must = fmt.Sprintf(templateUrl,"10")
	if s != must {
		t.Errorf("Expected %q, got %q",must, s)
	}

	s = formatUrl(">>",1,-1,templateUrl);
	must = ""
	if s != must {
		t.Errorf("Expected %q, got %q",must, s)
	}
	/// ...

	s = formatUrl(etc,1,10,templateUrl);
	must = ""
	if s != must {
		t.Errorf("Expected %q, got %q",must, s)
	}

	/// >
	s = formatUrl(">",10,10,templateUrl);
	must = fmt.Sprintf(templateUrl,"10")
	if s != must {
		t.Errorf("Expected %q, got %q",must, s)
	}

	s = formatUrl(">",1,10,templateUrl);
	must = fmt.Sprintf(templateUrl,"2")
	if s != must {
		t.Errorf("Expected %q, got %q",must, s)
	}



	s = formatUrl(">",0,10,templateUrl);
	must = ""
	if s != must {
		t.Errorf("Expected %q, got %q",must, s)
	}


	s = formatUrl("fail",0,10,templateUrl);
	must = ""
	if s != must{
		t.Errorf("Expected %q, got %q",must, s)
	}


	s = formatUrl("0",0,1,templateUrl);
	must = fmt.Sprintf(templateUrl,"1")
	if s != must{
		t.Errorf("Expected %q, got %q",must, s)
	}
}

func TestEmptyOrIncorrect(t *testing.T) {
	var r []*Pag

	testMustNotZero := func() {
		if r == nil {
			t.Error("It should not be nil")
		}
		if len(r)==0 {
			t.Error("It should not be empty, got", r)
		}
		if len(r) < 4 {
			t.Error("The length should be >=4 ", r)
		}

	}

	r = PagesArray(0, 0,templateUrl)
	testMustNotZero();

	r=PagesArray(0, 3,templateUrl);
	testMustNotZero();

	r=PagesArray(4, 0,templateUrl);
	testMustNotZero();

	r=PagesArray(10, 3,templateUrl);
	testMustNotZero();

	r=PagesArray(-23, -3,templateUrl);
	testMustNotZero();


}

func TestPriorNextOnBorder(t *testing.T) {
	var r []*Pag
	var pn = []string{"<<","<",">",">>"}

	r = PagesArray(0, 0,templateUrl)
	for i:=0; i<4; i++ {
		if r[i].Name != pn[i] {
			t.Errorf("Must be %q,got %q", pn[i],r[i].Name)
		}
	}

	var surl string
	r = PagesArray(1, 1,templateUrl)
	for i:=0; i<4; i++ {
		if surl="?page=1"; r[i].Url!=surl {
			t.Errorf("Must be %q,got %q",surl,r[i].Url)
		}
	}

	r = PagesArray(15, 100,templateUrl)
	if surl="?page=1"; r[0].Url!=surl {
		t.Errorf("Must be %q,got %q",surl,r[0].Url)
	}

	if surl="?page=100"; r[len(r)-1].Url!=surl {
		t.Errorf("Must be %q,got %q",surl,r[len(r)-1].Url)
	}
	if surl="?page=14"; r[1].Url!=surl {
		t.Errorf("Must be %q,got %q",surl,r[1].Url)
	}
	if surl="?page=16"; r[len(r)-2].Url!=surl {
		t.Errorf("Must be %q,got %q",surl,r[len(r)-2].Url , r)
	}

	r = PagesArray(98, 100,templateUrl)
	if surl="?page=100"; r[len(r)-1].Url!=surl {
		t.Errorf("Must be %q,got %q",surl,r[len(r)-1].Url)
	}

	if surl="?page=99"; r[len(r)-2].Url!=surl {
		t.Errorf("Must be %q,got %q",surl,r[len(r)-2].Url)
	}

	r = PagesArray(99, 100,templateUrl)
	if surl="?page=100"; r[len(r)-1].Url!=surl {
		t.Errorf("Must be %q,got %q",surl,r[len(r)-1].Url)
	}

	if surl="?page=100"; r[len(r)-2].Url!=surl {
		t.Errorf("Must be %q,got %q",surl,r[len(r)-2].Url)
	}

}

func TestPagesOrdered(t * testing.T) {
	var r []*Pag
	r = PagesArray(15, 100,templateUrl)
	/// << < 14 15 16 ... 98 99 100 > >>
	/// 0  1  2  3  4  5   6  7  8  9 10
	idxs := []string{"<<", "<", "14", "15", "16", "...", "98", "99" ,"100",">",">>"}
	for i, pag := range r {
		if idxs[i] !=  pag.Name {
			t.Fatalf("Position:%d, Got: %q, Must:%q",i,pag.Name ,idxs[i])
		}
	}
}


func TestPageArrayConcrete(t *testing.T) {
	var r []*Pag
	TestArr := func(arrValues []string, arrGot []*Pag) {
		if (len(arrValues)==0) {
			t.Fatalf("Must have items", arrGot)
		}
		if len(arrValues)+4!=len(arrGot) {
			t.Fatalf("Size %v = %d,but must be %d (%v)", arrGot, len(arrGot), len(arrValues)+4, arrValues)
		}

		for i := 0; i<len(arrValues); i++ {
			if arrValues[i]==etc {
				if arrGot[i+2].Url !="" {
					t.Errorf("Got url %q for %q, must be empty", arrGot[i+2].Url, arrValues[i])
				}
			} else {
				if arrGot[i+2].Url=="" {
					t.Fatalf("Got %v, index=%d, key %q is empty", arrGot, i, arrValues[i])
				}

				if sfmt := fmt.Sprintf(templateUrl, arrValues[i]); arrGot[i+2].Url!=sfmt {
					t.Errorf("Got url %q for %q, must be %q", arrGot[i].Url, arrValues[i], sfmt)
				}
			}
		}
	}

	r = PagesArray(1, 1,templateUrl)
	TestArr([]string{"1"}, r)

	r = PagesArray(1, 2,templateUrl)
	TestArr([]string{"1", "2"}, r)

	r = PagesArray(2, 2,templateUrl)
	TestArr([]string{"1", "2"}, r)


	r = PagesArray(2, 3,templateUrl)
	TestArr([]string{"1", "2", "3"}, r)

	r = PagesArray(1, 4,templateUrl)
	TestArr([]string{"1", "2", "3", "4"}, r)

	r = PagesArray(1, 5,templateUrl)
	TestArr([]string{"1", "2", "3", "4", "5"}, r)

	r = PagesArray(1, 6,templateUrl)
	TestArr([]string{"1", "2", "3", "4", "5", "6"}, r)

	r = PagesArray(1, 7,templateUrl)
	TestArr([]string{"1", "2", "3", "...", "5", "6", "7"}, r)

	r = PagesArray(1, 8,templateUrl)
	TestArr([]string{"1", "2", "3", "...", "6", "7", "8"}, r)

	r = PagesArray(1, 14,templateUrl)
	TestArr([]string{"1", "2", "3", "...", "12", "13", "14"}, r)

	r = PagesArray(1, 1001,templateUrl)
	TestArr([]string{"1", "2", "3", "...", "999", "1000", "1001"}, r)

	r = PagesArray(2, 1001,templateUrl)
	TestArr([]string{"1", "2", "3", "...", "999", "1000", "1001"}, r)

	r = PagesArray(3, 1001,templateUrl)
	TestArr([]string{"2", "3", "4", "...", "999", "1000", "1001"}, r)



	r = PagesArray(4, 1002,templateUrl)
	TestArr([]string{"3", "4", "5", "...", "1000", "1001", "1002"}, r)

	r = PagesArray(1001, 1002,templateUrl)
	TestArr([]string{"1", "2", "3", "...", "1000", "1001", "1002"}, r)
}
