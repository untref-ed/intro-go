# Introducción a Go

## ¿A quien va dirigido este tutorial?

Este tutorial está pensado para las personas que cursaron la materia Algoritmos
y Programación 2 utilizando el lenguaje de programación Java.

## Links útiles

- [Página principal](https://go.dev/)
- [Descargar Go](https://go.dev/dl/)
- [Taller de introducción a Go de AyP2](https://github.com/untref-ayp2/taller-GO)
- [Tour interactivo de Go](https://go.dev/tour/)
- [Documentación basada en ejemplos](https://gobyexample.com/)
- [Go FAQ](https://go.dev/doc/faq)
- [Go Playground](https://go.dev/play/) (entorno online para ejecutar código en
  Go)

## Similitudes entre Java y Go

- Son lenguajes compilados y con chequeo estático de tipos.
- Implementan un recolector de basura (_Garbage Collector_).

## Principales diferencias entre Java y Go

- Go no es orientado a objetos.
- Go utiliza punteros de forma directa.
- Go permite devolver multiples valores desde una función o un método.
- Go no tiene excepciones.
- Go tiene interfaces pero funcionan de manera distinta que en Java.

## Conceptos importantes que necesitamos entender de Go

### Estructura de un programa en Go

Utilizando el clásico ejemplo de un hola mundo, vamos a mostrar cómo se organiza
el código en Go.

El código se organiza en paquetes. Por ejemplo, podemos definir un paquete
llamado `saludo`, con una función que se exporta llamada `Saludar`.

[`saludo/saludo.go`](./holamundo/saludo/saludo.go)

```go
package saludo

import "fmt"

func Saludar() {
    fmt.Println("¡Hola mundo!")
}
```

Como en muchos lenguajes de programación, el punto de entrada a un programa es
por medio de una función `main`, como podemos ver a continuación, donde nuestro
programa hace uso del paquete `saludo`:

[`main.go`](./holamundo/main.go)

```go
package main

import "holamundo/saludo"

func main() {
    saludo.Saludar()
}
```

---

### Declaración de variables

Existen varias formas de declarar una variable en Go. Primero es utilizando la
palabra clave `var`, donde debemos especificar el nombre de la variable y su
tipo.

```go
var edad int
```

Opcionalmente podríamos inicializar esa variable en la misma línea como sucede
en muchos lenguajes de programación.

```go
var edad int = 42
```

Otra forma de declarar una variable es utilizando la notación de asignación
corta (`:=`), que determina el tipo de la variable de forma implícita. Por lo
que podemos escribir:

```go
edad := 42
```

Tipos básicos de datos en Go:

- `bool`
- `string`
- `int`, `int8`, `int16`, `int32`, `int64`
- `uint`, `uint8`, `uint16`, `uint32`, `uint64`
- `float32`, `float64`
- `complex64`, `complex128`
- `byte` (alias de `uint8`)
- `rune` (alias de `int32`, representa una posición en código Unicode)
- `uintptr` (tipo utilizado para guardar una dirección de puntero)

Go es un lenguaje fuertemente tipado. A diferencia de Java, Go no hace
conversión automática de tipos por lo que cada vez que necesitemos pasar de un
tipo a otro debemos realizar un casteo explicito.

```go
var edad int32 = 42
var edad64 int64 = edad
```

Este código data como resultado el siguiente error.

```text
cannot use edad (variable of type int32) as int64 value in variable declaration
```

En cambio, debemos realizar el casteo explicito para indicar al compilador de Go
que realmente queremos asignar un valor de tipo `int32` a una variable de tipo
`int64`.

```go
var edad int32 = 42
var edad64 int64 = int64(edad)
```

---

### Control de flujo en Go

Los bloques condicionales `if`/`else` se utilizan de forma muy similar a como
estamos acostumbrado en Java (con la salvedad que los paréntesis no son
necesarios en las condiciones).

```go
if num < 0 {
    fmt.Println(num, "es negativo")
} else if num < 10 {
    fmt.Println(num, "tiene 1 dígito")
} else {
    fmt.Println(num, "tiene múltiples dígitos")
}
```

A diferencia de Java, Go cuenta con sólo una instrucción de iteración, el `for`.
Pero este se puede utilizar de diferentes formas.

```go
i := 1

for i <= 3 {
    fmt.Println(i)
    i = i + 1
}
```

De forma análoga a un `while` como lo vimos en Java, cuando `for` solo recibe
una condición, va a ejecutar el bloque de código "mientras" la condición sea
verdadera.

```go
for j := 7; j <= 9; j++ {
    fmt.Println(j)
}
```

También puede utilizarse de la forma clásica, indicando la inicialización, la
condición y la operación luego de cada iteración.

```go
for {
    fmt.Println("loop")
    break
}
```

Si `for` no recibe condición, se comporta de la misma forma que `while(true)`.

```go
for n := 0; n <= 5; n++ {
    if n%2 == 0 {
        continue
    }
    fmt.Println(n)
}
```

También existen las instrucciones `break` y `continue` para alterar la ejecución
de las iteraciones.

---

### Slices

Los **slices** en Go son similares a `List` en Java. En Go también existen los
arrays de tamaño fijo.

```go
s := []int{1, 2, 3}
s = append(v1, 4)
```

Un slice es como un arreglo dinámico, al que le podemos agregar elementos luego
de haberlos declarado (a diferencia de los arrays básicos, con tamaño fijo).

```go
len(s)
```

Para conocer el tamaño actual de un slice, deberos utilizar la función
_built-in_ de Go `len`. En este caso `s` tiene una longitud de `4`.

```go
s[1]
```

Para acceder a los elementos de la lista lo hacemos por medio de indices (cómo
lo haríamos con un array convencional).

```go
s[1:3]
```

El nombre _slice_ está dado por la capacidad de estas estructuras de devolver
"tajadas", mediante el operador _slice_ (el operador `[:]`, si es un poco
confuso). En nuestro ejemplo anterior: `s[1:3]` devuelve un slice `[]int{2, 3}`
que es el rango de valores entre la posición 1 hasta la 3 (sin incluir esta
última).

Un array puede ser convertido en un _slice_ utilizando el operador de _slicing_
sobre el array: `arr[:]`.

---

### Maps

Los **maps** están implementados como `HashMap` y permiten buscar un valor por
un clave de forma eficiente. Para crear un map vacío, hacemos lo siguiente:

```go
m := make(map[string]int)

m["Juan"] = 1337

v, ok := m["Juan"]
```

Los valores se asignan y se leen por medio de una notación similar a la de los
arrays/slices. Es importante destacar que leer el contenido para una determinada
clave, devuelve 2 valores. En primer lugar el valor almacenado y en segundo
lugar un booleano que indica si la clave existe en el `map` o no.

Si el `map` no contiene la clave solicitada, entonces se devuelve el "valor
nulo" del tipo de los valores.

```go
m["Jose"] = 14
m["Maria"] = 0

v1, ok := m["Jose"]
fmt.Println("v1:", v1, "ok:", ok)

v2, ok := m["Maria"]
fmt.Println("v2:", v2, "ok:", ok)

vNot, ok := m["Pedro"]
fmt.Println("vNot:", vNot, "ok:", ok)
```

Ejecutar este código genera la siguiente salida:

```text
v1: 14 ok: true
v2: 0 ok: true
vNot: 0 ok: false
```

---

### Manejo de errores

Si bien Go no implementa excepciones como una construcción propia del lenguaje,
es posible hacer un control de errores por medio del módulo `errors`.

Por eso, en Go es común que los errores sean devueltos "normalmente" desde una
función y quien invocó dicha función es responsable de manejar un posible error.

```go
func test(input int) error {
    if input < 0 {
        return errors.New("less than zero")
    }
    return nil
}

func main() {
    err := test(-1)
    if err != nil {
        fmt.Print(err)
    }
}
```

En este ejemplo, podemos observar la forma típica en la que se maneja un error
en Go. Van a encontrar `if err != nil` en multiples lugares de un programa de
Go.

Go tiene una forma de "lanzar" errores por medio de la función `panic`, pero
esta debería ser reservada para casos muy extremos, ya que no es posible
recuperar el programa luego de esa invocación.

---

### Structs

En Go las _structs_ son colecciones de campos, podríamos pensar una `struct`
cómo una clase que solo declara atributos (los campos).

```go
type Persona struct {
    nombre, apellido string
    edad             uint
    direccion        Direccion
}
```

Para acceder a un campo de una estructura, se usa la notación de punto, como lo
hacemos para acceder a un atributo en Java.

```go
var p1 Persona
p1.nombre = "Marcelo"
p1.edad = 27

fmt.Println(p1.edad)
```

También podemos declarar una variable de tipo `struct` de forma literal:

```go
p2 := Persona{nombre: "Laura", apellido: "Medina", edad: 25}

fmt.Println(p2)
```

---

### Punteros

Un puntero es una variable que almacena una dirección de memoria a otra variable
de un determinado tipo.

Si asignamos un `int`, un `struct` o un `array`, se copia el contenido del
elemento. Para lograr el mismo efecto que con las referencias a variables en
Java, Go usa punteros.

Para cualquier tipo `T`, existe un correspondiente tipo puntero `*T`, que
determina un puntero a un valor de tipo `T`.

Si declaramos una variable de tipo `int`, podemos crear un puntero que apunte a
esa variable, este puntero será de tipo `*int`:

```go
var a int = 7
var pa *int = &a

fmt.Println("pa:", pa)
fmt.Println("*pa:", *pa)
```

La salida de este programa es:

```plain
pa: 0x14000098010
*pa: 7
```

Los punteros son tanto poderosos como simples, pero es posiblemente uno de los
temas más difíciles de asimilar. Hablemos un poco de sintaxis, cuando trabajamos
con punteros utilizamos 2 operadores, por un lado el operador de "desreferencia"
que es `*` y por otro lado el operador `&` que devuelve la posición de memoria
de una variable.

Cómo podemos ver en nuestro código estamos obteniendo la dirección de memoria de
la variable `a` (es decir su puntero) y lo almacenamos en la variable `pa` que
es de tipo `*int` (no confundir el operador `*` con el asterisco que se agrega
delante del nombre de un tipo para indicar que en realidad es el **puntero** a
un tipo).

El término "desreferenciar" significa tomar la posición de memoria y hacer
referencia al contenido de esa posición de memoria, no a la dirección en sí. Por
eso, cuando hacemos `*pa` estamos refiriendo al espacio de memoria donde fué
definida la variable `a`.

En Java, dependiendo el tipo de dato, los valores eran pasados por referencia o
por valor. En Go, nosotros debemos explicitar si vamos a recibir un valor por
referencia o por valor, al momento de definir los argumentos de nuestras
funciones.

---

### Funciones

Java no cuenta con el concepto de función como un componente independiente. En
Go las funciones pueden ser independientes del contexto y no es necesario que
sean definidas como un método estático de una clase para poder empaquetar
comportamiento. Las funciones en Go son "ciudadanos de primera clase", es decir
que una función es un valor, es decir que puedo almacenar una función en una
variable y pasarla como argumento de otra función o devolver una función como
resultado de alguna función.

Una función puede tomar cero o más argumentos, y puede devolver cero o más
valores.

```go
func sumar(x int, y int) int {
    return x + y
}
```

En este ejemplo la función `sumar` toma 2 parámetros de tipo `int`. Si
quisiéramos generar algo similar en Java deberíamos declarar una clase con
métodos estáticos:

```java
public class Matematica {
    public static int sumar(int a, int b) {
        return a + b;
    }
}
```

Cabe destacar, que a diferencia de Java, el valor de retorno en la declaración
de la función viene después de los paréntesis.

En Go es posible, devolver múltiples valores de una función, lo cual es muy
utilizado para el reporte de errores o la notificación del resultado de un
computo, cómo cuando se busca un valor por una clave en un `map`. Por ejemplo:

```go
func division_segura(dividendo, divisor float32) (float32, error) {
    if divisor == 0.0 {
        return 0.0, errors.New("división por cero")
    }

    return dividendo / divisor, nil
}
```

Cuando devolvemos múltiples valores los tipos devueltos se encierra entre
paréntesis, en el orden correspondiente.

---

### Métodos

Go no tiene clases como Java, sin embargo permite definir métodos sobre ciertos
tipos.

Un método es una función con un argumento especial **receptor** . El
**receptor** aparece en su propia lista de argumentos entre la palabra clave
`func` y el nombre del método.

```go
type Vector struct {
    X, Y float64
}

func (v Vector) Modulo() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vector) Escalar(factor float64) {
    v.X *= factor
    v.Y *= factor
}
```

En este ejemplo, el método `Modulo` tiene un receptor de tipo `Vector` llamado
`v`. Y el método `Escalar` recibe un puntero a `Vector`, ya que en este contexto
es necesario contar con la referencia a la variable apuntada, ya que este método
modifica el "estado" del receptor.

---

### ~~Public/Privado~~ Exportado/No exportado

En Go, la visibilidad y el acceso a los componentes definidos en un módulo está
dado por una convención de nombres. Si los nombres dados a los distintos
elementos en un paquete (variables, funciones, estructuras, etc.) comienzan con
una letra mayúsculas, entonces ese elemento se considera exportado y por lo
tanto puede ser considerado como parte de la "interfaz pública" del paquete (en
términos de Java).

```go
package ejemplo

import "fmt"

var datoProtegido string = "mi contraseña"

var DatoPublico float32 = 3.1415926

func HolaMundo() {
    fmt.Println("¡Hola Mundo!")
}
```

Esto también aplica a nivel de los campos de una estructura:

```go
package seguridad

type Boveda struct {
    Id   int
    dato string
}
```

En este ejemplo, `Id` será un dato accesible desde cualquier punto en el
programa que importe este paquete. Pero `dato` solo será accesible desde dentro
del mismo paquete. Para modificar este valor podemos definir métodos que operen
sobre estos campos no exportados.

```go
package seguridad

func (b *Boveda) actualizarDato(dato string) {
    if dato {
        b.dato = dato
    }
}
```

---

### Interfaces

Como mencionamos anteriormente, en Go existe el concepto de interfaces pero
funcionan de forma algo diferente a como lo hacen en Java.

Un tipo `interface` se define como una conjunto de firmas de método. Un valor de
ese tipo de interfaz, puede contener a cualquier valor que implemente (todos)
esos métodos.

```go
type Caminante interface {
    Avanzar(pasos int)
    Girar(grados float32)
}
```

Luego si un tipo implementa todos esos métodos:

```go
func (p *Persona) Avanzar(pasos int) { /* ... */ }

func (p *Persona) Girar(grados float32) { /* ... */ }dele
```

Podemos pasar como parámetro una variable de tipo `Persona` siempre se espere un
argumento de tipo `Caminante`.

```go
func RealizarRecorrido(caminante Caminante) { /* ... */ }

p := Persona{}

RealizarRecorrido(p)
```

## ¿Orientación a objetos?

De [FAQ: Is Go an object-oriented language?](https://go.dev/doc/faq#Is_Go_an_object-oriented_language):

> Si y no. Aunque Go tiene tipos y métodos y permite un estilo de programación
> orientado a objetos, no existe una jerarquía de tipos. El concepto de
> "interfaz" en Go proporciona un enfoque diferente que creemos que es fácil de
> usar y, en cierto modo, más general. También hay formas de incrustar tipos en
> otros tipos para proporcionar algo análogo, pero no idéntico, a la creación de
> subclases. Además, los métodos en Go son más generales que en C++ o Java: se
> pueden definir para cualquier tipo de datos, incluso tipos integrados, como
> enteros "sin caja". No están restringidos a estructuras (clases).
>
> Además, la falta de una jerarquía de tipos hace que los "objetos" en Go
> parezcan mucho más ligeros que en lenguajes como C++ o Java.

Veamos como poder definir un component que puede tener estado y comportamiento
asociado.

[`stack/stack.go`](./tda/stack/stack.go)

```go
package stack

import "errors"

type Stack struct {
    data []string
}

// Push agrega `x` en el tope de la pila.
func (s *Stack) Push(x string) {
    s.data = append(s.data, x)
}

// Pop remueve y retorna el valor en el tope de la pila.
// Devuelve un `error` si la pila está vacía.
func (s *Stack) Pop() (string, error) {
    if s.Size() == 0 {
        return "", errors.New("empty stack")
    }

    x := s.data[s.Size()-1]
    s.data = s.data[:s.Size()-1]
    return x, nil
}

// Size devuelve el número de elementos en la pila.
func (s *Stack) Size() int {
    return len(s.data)
}
```

Por un lado definimos una estructura la cual tiene "campos" (similar a lo que
llamamos atributos de una clase en Java). Luego, vamos a definir métodos que
tienen como receptor al tipo struct que creamos, de esta forma estos métodos
tendrán acceso a modificar los campos que no son exportados (cuyos nombres
comienzan con minúsculas).

[`main.go`](./tda/main.go)

```go
package main

import (
 "fmt"
 "tda/stack"
)

func main() {
 s := stack.Stack{}

 s.Push("world!")
 s.Push("Hello, ")

 for s.Size() > 0 {
  if x, err := s.Pop(); err == nil {
   fmt.Print(x)
  }
 }

 fmt.Println()
}
```

En nuestro caso, para crear una variable de tipo `Stack` podemos utilizar la
forma literal de declarar una estructura. También es común crear un método
llamado `NewStack` que sirve como constructor y que devuelva un puntero a un
variable de tipo `Stack`, el cual puede recibir argumentos y asignar esos
valores a los campos de la estructura.

Una vez creada la variable, podemos usar la notación de punto, como
acostumbramos a hacerlo al invocar métodos de un objeto.
