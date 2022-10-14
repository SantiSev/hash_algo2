package diccionario

import (
	"diccionario/lista"
	"fmt"
)

type hashDato[K comparable, V any] struct {
	clave    K
	variable V
}

type hashMap[K comparable, V any] struct {
	hashArray []lista.Lista[hashDato[K, V]] //todo O(1) tonces todo bien 8)
	longitud  int
}

func (h hashMap[K, V]) convertir(T any) int {
	dato := convertirABytes[K](T)
	fmt.Println(dato)
	index := h.sdbmHash(dato)
	fmt.Println(index)
	return index
}

func convertirABytes[K comparable](clave any) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func (h hashMap[K, V]) sdbmHash(data []byte) int {
	// documentacion: https://www.programmingalgorithms.com/algorithm/sdbm-hash/c/
	var hash uint64

	for _, b := range data {
		hash = uint64(b) + (hash << 6) + (hash << 16) - hash
	}

	return int(hash) % h.longitud
}

func (h hashMap[K, V]) Guardar(clave K, dato V) {
	//tu vieja
}

func (h hashMap[K, V]) Pertenece(clave K) bool {
	//TODO implement me
	panic("implement me")
}

func (h hashMap[K, V]) Obtener(clave K) V {
	//TODO implement me
	panic("implement me")
}

func (h hashMap[K, V]) Borrar(clave K) V {
	//TODO implement me
	panic("implement me")
}

func (h hashMap[K, V]) Cantidad() int {
	//TODO implement me
	panic("implement me")
}

func (h hashMap[K, V]) Iterar(f func(clave K, dato V) bool) {
	//TODO implement me
	panic("implement me")
}

func (h hashMap[K, V]) Iterador() IterDiccionario[K, V] {
	//TODO implement me
	panic("implement me")
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {

	return new(hashMap[K, V])
}

/*
		    .  ."|
   /| /  |  _.----._
  . |/  |.-"        ".  /|
 /                    \/ |__
|           _.-"""/        /
|       _.-"     /."|     /
 ".__.-"         "  |     \
    |              |       |
    /_      _.._   | ___  /
  ."  ""-.-"    ". |/.-.\/
 |    0  |    0  |     / |
 \      /\_     _/    "_/
  "._ _/   "---"       |
  /"""                 |
  \__.--                |_
    )          .        | ".
   /        _.-"\        |  ".
  /     _.-"             |    ".
 (_ _.-|                  |     |"-._.
   "    "--.             .J     _.-'
           /\        _.-" | _.-'
          /  \__..--"   _.-'
         /   |      _.-'
        /| /\|  _.-'
       / |/ _.-'
      /|_.-'
    _.-'
*/
