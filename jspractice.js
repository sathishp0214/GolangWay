console.log("hello")

var i = 20
console.log(i)

var g = 'sat' + "fgf"
console.log(i, g)
g = ""

if (g != "") {
    console.log('String not null---', g)
} else {
    console.log('String is null---', g)
}

for (i = 0; i < 10; i++) {
    console.log("for loop", i)
}

function sum(a, b) {
    return a + b
}

console.log(sum(10, 10))

x = [1, 2, 3, 4]
x.push(10)
console.log(x)
x.pop()
console.log(x)
x.push("sathish")
console.log(x)

for (i = 0; i < x.length; i++) {
    console.log(x[i])
}

