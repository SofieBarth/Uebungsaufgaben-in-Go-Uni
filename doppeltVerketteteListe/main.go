package main

import "fmt"

func main() {

	var listenElement1 listenElement
	var listenElement2 listenElement
	var listenElement3 listenElement
	var liste doppeltVerketteteListe

	listenElement1 = listenElement{
		wert:       5,
		vorheriges: nil,
		nächstes:   &listenElement2,
	}
	listenElement2 = listenElement{
		wert:       6,
		vorheriges: &listenElement1,
		nächstes:   &listenElement3,
	}
	listenElement3 = listenElement{
		wert:       7,
		vorheriges: &listenElement2,
		nächstes:   nil,
	}
	liste = doppeltVerketteteListe{
		anfang: &listenElement1,
		ende:   &listenElement3,
	}

	einfügenAmAnfang(&liste, 4)
	einfügenAmEnde(&liste, 8)
	//einfügenAnPosition(&liste, 9, 2)
	//fmt.Printf("%+v\n", gibElementAnPosition(&liste, 3))
	//listeAusgeben(&liste)
	//entferneAnfang(&liste)
	//entferneEnde(&liste)
	//entferneAnPosition(&liste, 2)
	listeAusgeben(&liste)
	//fmt.Println(gibPositionVonWert(&liste, 7))
}

// Verbund-Typ für ein Listen-Element
type listenElement struct {
	wert       int
	vorheriges *listenElement
	nächstes   *listenElement
}

// Verbund-Typ für eine Liste
type doppeltVerketteteListe struct {
	anfang, ende *listenElement
}

// Gibt alle Werte der Liste von Anfang bis Ende aus.
func listeAusgeben(liste *doppeltVerketteteListe) {
	element := liste.anfang
	fmt.Printf("-> ")
	for element != nil {
		fmt.Printf("%d -> ", element.wert)
		element = element.nächstes
	}
	fmt.Printf("nil\n")
}

// Gibt alle Werte der Liste von Ende bis Anfang aus.
func listeAusgebenRückwärts(liste *doppeltVerketteteListe) {
	element := liste.ende
	fmt.Printf("<- ")
	for element != nil {
		fmt.Printf("%d <- ", element.wert)
		element = element.vorheriges
	}
	fmt.Printf("nil\n")
}

// Liefert die Länge der Liste (d.h. die Anzahl ihrer Elemente) zurück.
func längeDerListe(liste *doppeltVerketteteListe) int {
	länge := 0
	element := liste.anfang
	for element != nil {
		länge++
		element = element.nächstes
	}
	return länge
}

// Fügt ein Element mit Wert neuerWert an den Anfang der Liste ein,
// d.h. vor das bisherige Anfangs-Element.
func einfügenAmAnfang(liste *doppeltVerketteteListe, neuerWert int) {
	neuesElement := &listenElement{
		wert:       neuerWert,
		vorheriges: nil,
		nächstes:   liste.anfang,
	}
	if liste.anfang == nil {
		liste.anfang = neuesElement
		liste.ende = neuesElement
		return
	}
	liste.anfang.vorheriges = neuesElement
	liste.anfang = neuesElement
}

// Fügt ein Element mit Wert neuerWert ans Ende der Liste an,
// d.h. hinter das bisherige End-Element.
func einfügenAmEnde(liste *doppeltVerketteteListe, neuerWert int) {
	neuesElement := &listenElement{
		wert:       neuerWert,
		vorheriges: liste.ende,
		nächstes:   nil,
	}
	if liste.ende == nil {
		liste.ende = neuesElement
		liste.anfang = neuesElement
		return
	}
	liste.ende.nächstes = neuesElement
	liste.ende = neuesElement
}

// Fügt ein neues Element mit Wert neuerWert an der Position pos eines bereits existierenden Elements
// in die Liste ein, wobei das existierende Element nach rechts verschoben wird.
func einfügenAnPosition(liste *doppeltVerketteteListe, neuerWert int, pos int) {
	if liste.anfang == nil {
		einfügenAmAnfang(liste, neuerWert)
		return
	}
	if pos > längeDerListe(liste) || pos < 0 {
		fmt.Println("Position nicht in der Liste enthalten")
		return
	}
	element := liste.anfang
	for i := 0; i < pos; i++ {
		element = element.nächstes
	}
	if element.vorheriges == nil {
		einfügenAmAnfang(liste, neuerWert)
		return
	}
	if pos == längeDerListe(liste) {
		einfügenAmEnde(liste, neuerWert)
		return
	}
	neuesElement := &listenElement{
		wert:       neuerWert,
		vorheriges: element.vorheriges,
		nächstes:   element,
	}
	element.vorheriges.nächstes = neuesElement
	element.vorheriges = neuesElement
}

// Liefert das Element an Position pos der Liste zurück.
func gibElementAnPosition(liste *doppeltVerketteteListe, pos int) *listenElement {
	if liste.anfang == nil {
		return nil
	}
	if pos >= längeDerListe(liste) || pos < 0 {
		fmt.Println("Position nicht in der Liste enthalten")
		return nil
	}
	elementAnPos := liste.anfang
	for i := 0; i < pos; i++ {
		elementAnPos = elementAnPos.nächstes
	}
	return elementAnPos
}

// Liefert den Wert des Elements an Position pos der Liste zurück.
func gibWertAnPosition(liste *doppeltVerketteteListe, pos int) int {
	if pos < 0 || pos >= längeDerListe(liste) {
		fmt.Printf("FEHLER: Position %v ist unzulässig in einer Liste der Länge %v.\n", pos, längeDerListe(liste))
		return -1
	}
	return gibElementAnPosition(liste, pos).wert
}

// Entfernt das Element am Anfang der Liste.
func entferneAnfang(liste *doppeltVerketteteListe) {
	if liste.anfang == nil {
		return
	}
	element := liste.anfang.nächstes
	element.vorheriges = nil
	liste.anfang = element
}

func entferneEnde(liste *doppeltVerketteteListe) {
	if liste.anfang == nil {
		return
	}
	element := liste.ende.vorheriges
	element.nächstes = nil
	liste.ende = element
}

// Entfernt das Element an der Position pos aus der Liste.
func entferneAnPosition(liste *doppeltVerketteteListe, pos int) {
	if liste.anfang == nil {
		return
	}
	if pos >= längeDerListe(liste) || pos < 0 {
		fmt.Println("Position nicht in der Liste enthalten")
		return
	}
	element := liste.anfang
	for i := 0; i < pos; i++ {
		element = element.nächstes
	}
	if element == liste.anfang {
		entferneAnfang(liste)
		return
	}
	if element == liste.ende {
		entferneEnde(liste)
		return
	}
	element.vorheriges.nächstes = element.nächstes
	element.nächstes.vorheriges = element.vorheriges
}

// Liefert die Position des ersten Auftretens des Werts gesuchterWert in der Liste zurück,
// oder -1, wenn der Wert nicht in der Liste enthalten ist.
// Wenn gesuchterWert mehrfach in der Liste enthalten ist,
// wird also immer die Position des Wertes zurückgeliefert, der am weitesten links in der Liste steht.
func gibPositionVonWert(liste *doppeltVerketteteListe, gesuchterWert int) int {
	if liste.anfang == nil {
		return -1
	}
	element := liste.anfang
	position := 0

	for element.wert != gesuchterWert && position < längeDerListe(liste) {
		element = element.nächstes
		position++
	}
	if element.wert == gesuchterWert {
		return position
	}
	return -1
}

// Liefert das (am weitesten links stehende) Element der Liste zurück,
// das den Wert gesuchterWert hat,
// oder nil, wenn kein Element mit Wert gesuchterWert in der Liste enthalten ist.
// Wenn mehrere Elemente mit gesuchterWert in der Liste enthalten sind,
// wird also immer dasjenige Element davon zurückgeliefert, das am weitesten links in der Liste steht.
func gibElementMitWert(liste *doppeltVerketteteListe, gesuchterWert int) *listenElement {
	i := gibPositionVonWert(liste, gesuchterWert)
	if i == -1 {
		return nil
	} else {
		return gibElementAnPosition(liste, i)
	}
}

// true, wenn der Wert gesuchterWert in der Liste enthalten ist,
// ansonsten false.
func istWertInListe(liste *doppeltVerketteteListe, gesuchterWert int) bool {
	return gibPositionVonWert(liste, gesuchterWert) != -1
}
