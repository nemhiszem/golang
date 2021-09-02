package main

//	var i int = 42 declarálni változót a func-on kívül
/*	var (
	actorname string = "Mór"
	companion string = "Mór2"
	season int = 11
)*/
//	var blokk amiben a benne lévő elemek declaráltnak számítanak

func main() {
	//	var i int első legenerálom és csak késöbb kap értéket
	//	i = 42
	//	var j float32 = 27 második egyszerre genereálom és adok neki értéket
	//	k := 99 harmadik rövid ilyenkor a complier dönti el miylen type a változó
	//	fmt.Printf("%v, %T", k, k) kiiratási módszer %v a változó értéke a %T pedig a típusa
	/*var i int = 42
	fmt.Printf("%v, %T\n", i, i)

	var j float32
	j = float32(i)
	fmt.Printf("%v, %T\n", j, j) ezzel lehet egyik változó értékét másmilyenre konvertálni egy másik változóba*/

	/*var i int = 42
	fmt.Printf("%v, %T\n", i, i)

	var j string
	j = strconv.Itoa(i)
	fmt.Printf("%v, %T\n", j, j) intből stringbe konvertálás, úgy hogy ne ASCII-ban jelenjen meg*/

}
