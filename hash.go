package diccionario

import "diccionario/lista"

type hashDato[K comparable, V any] struct {
	clave    K
	variable V
}

type hash[K comparable, V any] struct {
	hashArray []lista.Lista[hashDato[K, V]] //todo O(1) tonces todo bien 8)
	longitud  int
}

func (h hash[K, V]) Guardar(clave K, dato V) {
	//TODO implement me
	panic("implement me")
}

func (h hash[K, V]) Pertenece(clave K) bool {
	//TODO implement me
	panic("implement me")
}

func (h hash[K, V]) Obtener(clave K) V {
	//TODO implement me
	panic("implement me")
}

func (h hash[K, V]) Borrar(clave K) V {
	//TODO implement me
	panic("implement me")
}

func (h hash[K, V]) Cantidad() int {
	//TODO implement me
	panic("implement me")
}

func (h hash[K, V]) Iterar(f func(clave K, dato V) bool) {
	//TODO implement me
	panic("implement me")
}

func (h hash[K, V]) Iterador() IterDiccionario[K, V] {
	//TODO implement me
	panic("implement me")
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {

	return new(hash[K, V])
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
