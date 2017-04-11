package convnet



import (

"math"
"math/rand"
"time"

//"fmt"

//"github.com/pquerna/ffjson/ffjson"
//	"github.com/pquerna/ffjson/ffjson"
"encoding/json"
//"io/ioutil"

//"strings"
"strconv"
//"os"
"sort"

//"reflect"
"reflect"


//	"sync"
"sync"



//"net"

"os"

"image/color"
"image/png"
"image"


	"fmt"
)

var undefined_f float64 = -2147483648.0
var undefined_i int = -2147483648
var undefined_s string = ""
var undefined_ystruct = ystruct{dim: undefined_i, val: undefined_f}
var global Global

var now = time.Now()

type obj map[string]interface{}

func filter_float(f float64) float64 {

//if math.IsInf(f, -1) {
//return -0.0
//}
//if math.IsInf(f, 1) {
//return 0.0
//}
//if math.IsNaN(f) {
//return 0.0
//}
return f
}



func SetLayerDefString (foo Layer_def, field string, value string) Layer_def{
	if field == "type" {
		foo.Type = value

	} else if field == "activation" {
		foo.activation = value
	} else if field == "tensor" {
		foo.tensor = value
	}
	return foo
}

func SetLayerDefInt (foo Layer_def, field string, value int64) Layer_def{


	if field == "filters" {
	} else if field == "sx" {
		foo.sx = int(value)
	} else if field == "stride" {
		foo.stride = int(value)
	} else if field == "sy" {
		foo.sy = int(value)
	} else if field == "pad" {
		foo.pad = int(value)
	} else if field == "n" {
		foo.n = int(value)
	} else if field == "group_size" {
		foo.group_size = int(value)
	} else if field == "num_classes" {
		foo.num_classes = int(value)
	} else if field == "num_neurons" {
		foo.num_neurons = int(value)
	} else if field == "group_size" {
		foo.group_size = int(value)
	} else if field == "in_sx" {
		foo.in_sx = int(value)
	} else if field == "in_depth" {
		foo.in_depth = int(value)
	} else if field == "in_sy" {
		foo.in_sy = int(value)
	} else if field == "Out_depth" {
		foo.Out_depth = int(value)
	} else if field == "depth" {
		foo.depth = int(value)
	} else if field == "Out_sx" {
		foo.Out_sx = int(value)
	} else if field == "Out_sy" {
		foo.Out_sy = int(value)

	}


	return foo
}


func SetLayerDefFloat (foo Layer_def, field string, value float64) Layer_def {
	if field == "alpha" {
		foo.alpha = value
	} else if field == "beta" {
		foo.beta = value
	} else if  field == "bias_pref" {
		foo.bias_pref = value
	} else if  field == "drop_prob" {
		foo.drop_prob = value
	} else if  field == "l1_decay_mul" {
		foo.l1_decay_mul = value
	} else if  field == "l2_decay_mul" {
		foo.l2_decay_mul = value
	}
	return foo
}

func SetTrainerDefFloat (foo Trainer_def, field string, value float64) Trainer_def {

	if field == "learning_rate" {
		foo.learning_rate = value
	} else if field == "l1_decay" {
		foo.l1_decay = value
	} else if field == "l2_decay" {
		foo.l2_decay = value
	} else if field == "momentum" {
		foo.momentum = value
	} else if field == "eps" {
		foo.eps = value
	} else if field == "ro" {
		foo.ro = value
	}

	return foo
}


func SetTrainerDefInt(foo Trainer_def, field string, value int) Trainer_def {

	if field == "batch_size" {
		foo.batch_size = value
	}

	return foo
}

func SetTrainerDefString(foo Trainer_def, field string, value string) Trainer_def {

	if field == "method" {
		foo.method = value
	}

	return foo
}


func get_float( f float64) float64 {


//if f == nil  {
//	return 0.0
//}


return filter_float(f)
}

func get_int( f int) int {


//if f == nil  {
//	return 0.0
//}


return (f)
}

//func float_arr_to_int_arr (f []float64) []int {
//
//	var i []int = []int{}
//	for x := 0; x < len(f); x++ {
//
//		i = append(i, int(f[x]));
//
//	}
//
//	return i
//
//}
//
//func dup_mfloat (f []float64) []float64 {
//
//	return (mfloat_arr_to_float_arr(f))
//
//}

//func mfloat_arr_to_int_arr (f []float64) []int {
//
//	var i []int = []int{}
//	for x := 0; x < len(f); x++ {
//
//		i = append(i, int((f[x])));
//
//	}
//
//	return i
//
//}

//func float_arr_to_mint_arr (f []float64) []int {
//
//	var i []int = []int{}
//	for x := 0; x < len(f); x++ {
//
//		i = append(i, set_mint(int(f[x])));
//
//	}
//
//	return i
//
//}


//func mfloat_arr_to_mint_arr (f []float64) []int {
//
//	var i []int = []int{}
//	for x := 0; x < len(f); x++ {
//
//		i = append(i, set_mint(int((f[x]))));
//
//	}
//
//	return i
//
//}

//func int_arr_to_float_arr (f []int) []float64 {
//
//	var i []float64 = []float64{}
//	for x := 0; x < len(f); x++ {
//
//		i = append(i, float64(f[x]));
//
//	}
//
//	return i
//
//}
//
//func mint_arr_to_float_arr (f []int) []float64 {
//
//	var i []float64 = []float64{}
//	for x := 0; x < len(f); x++ {
//
//		i = append(i, float64(get_int(f[x])));
//
//	}
//
//	return i
//
//}

//func int_arr_to_mfloat_arr (f []int) []float64 {
//
//	var i []float64 = []float64{}
//	for x := 0; x < len(f); x++ {
//
//		i = append(i, (float64(f[x])));
//
//	}
//
//	return i
//
//}


//func mint_arr_to_mfloat_arr (f []int) []float64 {
//
//	var i []float64 = []float64{}
//	for x := 0; x < len(f); x++ {
//
//		i = append(i, (float64(get_int(f[x]))));
//
//	}
//
//	return i
//
//}



//func int_to_mfloat (f int) float64 {
//
//return (float64(f))
//
//}
//
//func float_to_mint (f float64) int {
//
//	return set_mint(int(f))
//
//}

func set_mint (f int) int {

//f = (f)

return f
}




//func  set_mfloat(f float64) float64 {
//
//	//f = filter_float(f)
//
//	return f
//}


//func dup_int (f int) int {
//
//	tf := get_int(f)
//	return tf
//}


func dup_float (f float64) float64 {

tf := f
return tf
}




func new_layer_def() layer_def {

var def layer_def
def.stride = undefined_i
def.pad = undefined_i
def.n = undefined_i
def.k = undefined_i
def.alpha = undefined_f
def.beta  = undefined_f
def.Type = undefined_s
def.Out_sx = undefined_i
def.Out_sy  = undefined_i
def.depth  = undefined_i
def.num_classes  = undefined_i
def.num_neurons  = undefined_i
def.bias_pref  = undefined_f
def.activation = undefined_s
def.tensor = undefined_s
def.group_size  = undefined_i
def.drop_prob  = undefined_f
def.in_sx  = undefined_i
def.in_depth  = undefined_i
def.in_sy  = undefined_i
def.Out_depth  = undefined_i
def.l1_decay_mul  = undefined_f

def.l2_decay_mul  = undefined_f
def.filters  = undefined_i
def.sx = undefined_i
def.sy = undefined_i

return def

}

type layer struct {
switchx [][]float64
switchy [][]float64
sx []int
sy  []int
stride []int
pad []int
S_cache_ []*Vol
in_sx []int
in_depth []int
in_sy []int
Out_depth []int
Out_sx []int
Out_sy []int
n []int
k []int
alpha []float64
beta []float64
layer_type []string
num_inputs []int
in_act []*Vol
out_act []*Vol
drop_prob []float64
dropped [][]float64
num_neurons []int
filters [][]*Vol
l1_decay_mul []float64
l2_decay_mul []float64
bias_pref []float64
biases []*Vol
switches [][]float64
group_size []int
es [][]float64
}

type LSTMNet struct {
Model map[string]*Vol
Hidden []*Vol
Cell []*Vol
needs_backprop bool
Backprop []func()
//===
}


type Net struct {
Layers []*layer
layer layer
score []float64
trained_epochs float64
}

type Vol struct {
sx int
sy int
depth int
c float64
w []float64
Dw []float64
d int
N int
}

func (this *Vol) W() []float64 {
	return this.w
}

func (this *Vol) Set(n int, v float64)  {
	this.w[n] = v
}




type fold struct {
train_ix []int
test_ix []int
}


type trainer_def struct {
learning_rate float64
l1_decay float64
l2_decay float64
batch_size int;
method string // sgd/adagrad/adadelta/windowgrad
momentum float64
eps float64
ro float64
}

type Trainer_def trainer_def


type candidate struct {

acc []float64
accv float64
layer_defs []layer_def
trainer_def trainer_def
net *Net
trainer *Trainer

}

// // Utility fun
  func assert(condition bool, message string) {
    // from http://stackoverflow.com/questions/15313418/javascript-assert
    if (!condition) {
		if message =="" {
			 fmt.Println("Assertion failed");
		}
      fmt.Println(message);
    }
  }


 // Mat holds a matrix
  func (this *Vol) Mat(n int,d int) {
    // n is number of rows d is number of columns
    this.N = n;
    this.d = d;
    this.w = global.zeros(n * d);
    this.Dw = global.zeros(n * d);
  }



type candidates struct {
self []*candidate //= make([]*candidate, 0)
}

type  MagicNet  struct {


data []*Vol
labels []float64
train_ratio float64
num_folds int
num_candidates int
num_epochs int
ensemble_size int
batch_size_min int
batch_size_max int
l2_decay_min float64
l2_decay_max float64
learning_rate_min float64
learning_rate_max float64
momentum_min float64
momentum_max float64
neurons_min int
neurons_max int

folds  []*fold
candidates *candidates
evaluated_candidates *candidates
unique_labels []float64
iter int
foldix int








}


func  copy_temp_Vol(v *Vol, c float64) *Vol {
this := new(Vol)
this.sx = v.sx
this.sy = v.sy
this.depth = v.depth
this.c = c
n:= this.sx*this.sy*this.depth
if n==undefined_i { this.w = make([]float64, 0); this.Dw = make([]float64, 0); } else {
this.w = make( []float64, n);
this.Dw = make( []float64, n);
}

if c == undefined_f {



for i:=0;i<n;i++ {
r := global.gaussRandom() // my addition

this.w[i] = (r); // possibly make this a new random number




}
} else {
for i:=0;i<n;i++ {



this.w[i] = (c);



}
}
return this

}

func (this *Vol) get_grad (x int, y int, d int) float64 {

var ix = ((this.sx * y)+x)*this.depth+d;
return this.Dw[ix];
}


func (this *Vol) add_grad (x int, y int, d int, v float64) {


var ix = ((this.sx * y)+x)*this.depth+d;

//if ix < 0 {
//	ix*= -1 // maybe this is wrong but it stops the app from crashing
//}

this.Dw[ix] = ((this.Dw[ix]) + v);
}

func (this *Vol) set_grad(x int, y int, d int, v float64) {

var ix = ((this.sx * y)+x)*this.depth+d;
this.Dw[ix] = (v);
}




func (this *Vol) cloneAndZero() *Vol {

return  New_Vol(this.sx, this.sy, this.depth, 0.0)



}


func (this *Vol) clone() *Vol {



v := New_Vol(this.sx, this.sy, this.depth, this.c)



for x := 0; x < len(v.w); x++ {

v.w[x] = dup_float(this.w[x])
v.Dw[x] = dup_float(this.Dw[x])

}



return v
}

func copyVol(dest *Vol, src *Vol) {

dest.sx = src.sx
dest.sy = src.sy
dest.depth = src.depth
dest.c = src.c
dest.w  = src.w
dest.Dw = src.Dw
//dest.w = make([]float64, len(src.w))
//dest.Dw = make([]float64, len(src.Dw))
//for x := 0; x < len(src.w); x++ {
//	dest.w[x] = dup_float(src.w[x])
//	dest.w[x] = dup_float(src.Dw[x])
//}
}

func duplicate_Vol(src *Vol) *Vol {
var dest *Vol = new(Vol)


dest.sx = src.sx
dest.sy = src.sy
dest.depth = src.depth
dest.c = src.c
//dest.w  = src.w
//dest.Dw = src.Dw
dest.w = make([]float64, len(src.w))
dest.Dw = make([]float64, len(src.Dw))
for x := 0; x < len(src.w); x++ {
dest.w[x] = dup_float(src.w[x])
dest.w[x] = dup_float(src.Dw[x])
}


return dest

}

func (this *Vol) get (x int, y int, d int) float64 {

var ix=((this.sx * y)+x)*this.depth+d;
return this.w[ix];
}

func (this *Vol) set (x int, y int, d int, v float64) {
var ix=((this.sx * y)+x)*this.depth+d;
this.w[ix] = (v);
}



func New_Vol_from_array(a []float64, net *Net) *Vol {
this := new(Vol)

//// this is how you check if a variable is an array. Oh, Javascript :)
//if(Object.prototype.toString.call(sx) === '[object Array]') {
//// we were given a list in sx, assume 1D Volume and fill it up
//this.sx = 1;
//this.sy = 1;
//this.depth = sx.length;
//// we have to do the following copy because we want to use
//// fast typed arrays, not an ordinary javascript array
//this.w = global.zeros(this.depth);
//this.Dw = global.zeros(this.depth);
//for(var i=0;i<this.depth;i++) {
//this.w[i] = sx[i];
//}
//} else {
// we were given dimensions of the Vol
this.sx = net.Layers[0].Out_sx[0];






this.sy = net.Layers[0].Out_sy[0];
this.depth = net.Layers[0].Out_depth[0];
var n = this.sx*this.sy*this.depth;

this.w = global.zeros(n);
this.Dw = global.zeros(n);
for x:=0; x < len(a); x++ {
this.w[x] = (a[x])
}
return this
}


func  New_Vol (sx int, sy int, depth int, c float64) *Vol {


this := new(Vol)
//// this is how you check if a variable is an array. Oh, Javascript :)
//if(Object.prototype.toString.call(sx) === '[object Array]') {
//// we were given a list in sx, assume 1D Volume and fill it up
//this.sx = 1;
//this.sy = 1;
//this.depth = sx.length;
//// we have to do the following copy because we want to use
//// fast typed arrays, not an ordinary javascript array
//this.w = global.zeros(this.depth);
//this.Dw = global.zeros(this.depth);
//for(var i=0;i<this.depth;i++) {
//this.w[i] = sx[i];
//}
//} else {
// we were given dimensions of the Vol
this.sx = sx;
this.sy = sy;
this.depth = depth;
var n = sx*sy*depth;
this.w = global.zeros(n);
this.Dw = global.zeros(n);

//copy(this.Dw,this.w) // this might be a little stffup

if(c == undefined_f) {


// weight normalization is done to equalize the output
// variance of every neuron, otherwise neurons with a lot
// of incoming connections have outputs of larger variance
var scale = math.Sqrt(1.0/float64(sx*sy*depth));



////sem := thread_make(n)
//wg := sync.WaitGroup{}

//wg.Add(n)

for i:=0;i<n;i+=1 {
//go func(i int) {

//	defer wg.Done()
////go func(i int) {


t := global.randn(0.0, scale);

this.w[i] = (t)
////thread_signal(sem)
//}(i)
//}(i)
}

//wg.Wait()

////thread_wait(n, sem)












} else {

////sem := thread_make(n)


for i:=0;i<n;i++ {

//	//go func(i int) {

this.w[i] = (c);


////thread_signal(sem)
//}(i)




}

////thread_wait(n, sem)
}
//}




return this

}

type Global struct {

//layer layer
//depth int
//num_classes int
//num_neurons int
////in_act *Vol
//out_act *Vol
//
//
//
//layer_def layer_def
//
return_v bool
v_val float64
//num_inputs int
}
/* mutexes */
type empty struct {}
type semaphore chan empty

// acquire n resources
func (s semaphore) P(n int) {
e := empty{}
for i := 0; i < n; i++ {
s <- e
}
}

// release n resources
func (s semaphore) V(n int) {
for i := 0; i < n; i++ {
<-s
}
}

func (s semaphore) Lock() {
s.P(1)
}

func (s semaphore) Unlock() {
s.V(1)
}

/* signal-wait */

func (s semaphore) Signal() {
s.V(1)
}

func (s semaphore) Wait(n int) {
s.P(n)
}




func (this *Global) zeros(n int) []float64 {
if n==undefined_i { return make([]float64, 0); } // maybe change the 0 to a 1
//if(typeof ArrayBuffer === 'undefined') {
// lacking browser support






arr := make( []float64, n);


return  arr
//=== done
//	wg := sync.WaitGroup{}

//	wg.Add(n)
//	for i := 0; i < n; i++ {
//
//	//go func(i int) {
//	//	defer wg.Done()
//
//		arr[i] = (0.0)
//
//
//	//}(i)
//
//
//	}
//
//	//wg.Wait()
//
//
//
//
//
//
//	return arr;
//} else {
//return new Float64Array(n);
//}
}


func (this *Global) randn(mu float64, std float64) float64 {


return mu+this.gaussRandom()*std;
}

func (this *Global) gaussRandom () float64{

if(this.return_v) {
this.return_v = false;

return this.v_val;
}



source := rand.NewSource(time.Now().UnixNano())
generator := rand.New(source)
//rand.Seed(time.Now().UTC().UnixNano())
generator.Seed(time.Now().UTC().UnixNano())

return generator.NormFloat64() // maybe take this out

var u = (2*rand.Float64()-1) ; // maybe no / 2
var v = (2*rand.Float64()-1); // maybe no / 2



var r = u*u + v*v;

if(r == 0 || r > 1) {  return this.gaussRandom() };

var c = math.Sqrt(-2*math.Log(r)/r);
this.v_val = v*c; // cache this
this.return_v = true;
return u*c;
//ad
}


func (this *Net) getParamsAndGrads () []ParamsAndGrads {

// accumulate parameters and gradients for the entire network
var response []ParamsAndGrads = make([]ParamsAndGrads, 0);
var temp [][]ParamsAndGrads = make([][]ParamsAndGrads, len(this.Layers));
//sem := thread_make(len(this.Layers))
//fmt.Println(this.Layers[len(this.Layers)-1].layer_type)
wg := sync.WaitGroup{}
wg.Add(len(this.Layers))

for i:=0;i<len(this.Layers);i++ {

go func(i int) {

defer wg.Done()

var layer_reponse = this.Layers[i].getParamsAndGrads(i);
//
temp_response := make([]ParamsAndGrads, len(layer_reponse))

n := len(layer_reponse)

for j := 0; j < n; j++ {
temp_response[j] = (layer_reponse[j]);
}


//}(len(layer_reponse), layer_reponse)



temp[i] = temp_response

//thread_signal(sem)

}(i)
}

wg.Wait()
//thread_wait(len(this.Layers), sem)
for i:=0;i<len(this.Layers);i++ {

response = append(response, temp[i]...);

}
return response;

}
func (this *Net) getPrediction () float64 {

var S *layer = this.Layers[len(this.Layers)-1]; // softmax layer
var p []float64 = S.out_act[len(this.Layers)-1].w; // maybe this should be 0 not len(S.out_act)-1

//p := v.w

//println((p[0]))
//println(len(S.out_act)-1, "p", "getPrediction")
var maxv float64 = (p[0]);
var maxi float64 = 0.0;
for i:=1;i<len(p);i++ {
if((p[i]) > maxv) {
maxv = (p[i]); maxi = float64(i);
}
}




return maxi;


}

func Forward (this *Net, V interface{}) []float64 {


	return this.forward(V.(*Vol), false).w


}

func (this *Net) forward (V *Vol, is_training bool) *Vol {


//var V *Vol = duplicate_Vol(x)


///if is_training == false { is_training = false };
var act *Vol = this.Layers[0].forward(V, is_training, 0);



for i:=1;i<len(this.Layers);i++ {


//
act = this.Layers[i].forward(act, is_training, i);
//println(i)
//ln((this.Layers[i].filters[1][0].w[0]))

}

//x = act
return act;
}


func (this *Net) backward (y []float64, ys ystruct) float64 {
// maybe i dont need this function

var N = len(this.Layers);
var loss = this.Layers[N-1].backward(y, N-1, ys); // last layer assumed softmax
for i:=N-2;i>=0;i-- { // first layer assumed input
_ = this.Layers[i].backward([]float64{}, i, ys);
}
return loss;

}
func (this *layer) PoolLayer(opt layer_def, ln int) *layer {


this.append_layer();

// required
this.sx[ln] = opt.sx; // filter size
this.in_depth[ln] = opt.in_depth;
this.in_sx[ln] = opt.in_sx;
this.in_sy[ln] = opt.in_sy;


//println(opt.in_sx)

// optional
if opt.sy != undefined_i {
this.sy[ln] = opt.sy
} else {
this.sy[ln] = this.sx[ln]
}


if opt.stride != undefined_i {
this.stride[ln] = opt.stride
} else {
this.stride[ln] = 2
}


if opt.pad != undefined_i {
this.pad[ln] = opt.pad
} else {
this.pad[ln] = 0
}


// computed
this.Out_depth[ln] = this.in_depth[ln];
this.Out_sx[ln] = int(math.Floor(float64(this.in_sx[ln] + this.pad[ln] * 2 - this.sx[ln]) / float64(this.stride[ln]) + 1.0));
this.Out_sy[ln] = int(math.Floor(float64(this.in_sy[ln] + this.pad[ln] * 2 - this.sy[ln]) / float64(this.stride[ln]) + 1.0));
this.layer_type[ln] = "pool";
// store switches for x,y coordinates for where the max comes from, for each output neuron
n:= ((this.Out_sx[ln]*this.Out_sy[ln]*this.Out_depth[ln]))

//println(n)

this.switchx[ln] = global.zeros(n);
this.switchy[ln] = global.zeros(n);

return this
}




func (this *layer) ConvLayer(opt layer_def, ln int) *layer {
this.append_layer();


//println(opt.filters)
// required
this.Out_depth[ln] = opt.filters;
this.sx[ln] = opt.sx; // filter size. Should be odd if possible, it's cleaner.
this.in_depth[ln] = opt.in_depth;
this.in_sx[ln] = opt.in_sx;
this.in_sy[ln] = opt.in_sy;




// optional
if opt.sy != undefined_i {
this.sy[ln] = opt.sy

} else {
this.sy[ln] = this.sx[ln]

}
//this.sy[ln] = typeof opt.sy !== 'undefined' ? opt.sy : this.sx;

if opt.stride != undefined_i {
this.stride[ln] = opt.stride
} else {
this.stride[ln] = 1
}

//this.stride[ln] = typeof opt.stride !== 'undefined' ? opt.stride : 1; // stride at which we apply filters to input Volume


if opt.pad != undefined_i {
this.pad[ln] = opt.pad
} else {
this.pad[ln] = 0
}

//this.pad[ln] = typeof opt.pad !== 'undefined' ? opt.pad : 0; // amount of 0 padding to add around borders of input Volume

if opt.l1_decay_mul != undefined_f {
this.l1_decay_mul[ln] = opt.l1_decay_mul
} else {
this.l1_decay_mul[ln] = 0.0
}

//this.l1_decay_mul[ln] = typeof opt.l1_decay_mul !== 'undefined' ? opt.l1_decay_mul : 0.0;

if opt.l2_decay_mul != undefined_f {
this.l2_decay_mul[ln] = opt.l2_decay_mul
} else {
this.l2_decay_mul[ln] = 0.0
}

//this.l2_decay_mul[ln] = typeof opt.l2_decay_mul !== 'undefined' ? opt.l2_decay_mul : 1.0;

// computed
// note we are doing floor, so if the strided conVolution of the filter doesnt fit into the input
// Volume exactly, the output Volume will be trimmed and not contain the (incomplete) computed
// final application.
this.Out_sx[ln] = int(math.Floor((float64(this.in_sx[ln]) + float64(this.pad[ln]) * 2 - float64(this.sx[ln])) / float64(this.stride[ln]) + 1.0));
this.Out_sy[ln] = int(math.Floor((float64(this.in_sy[ln]) + float64(this.pad[ln]) * 2 - float64(this.sy[ln])) / float64(this.stride[ln]) + 1));
this.layer_type[ln] = "conv";

// initializations
var bias float64
if opt.bias_pref != undefined_f {
bias = opt.bias_pref
} else {
bias = 0.0
}


//var bias = typeof opt.bias_pref !== 'undefined' ? opt.bias_pref : 0.0;
this.filters[ln] = make([]*Vol, 0);
for i:=0;i<this.Out_depth[ln];i++ {
this.filters[ln]= append(this.filters[ln],New_Vol(this.sx[ln], this.sy[ln], this.in_depth[ln], undefined_f)); }
this.biases[ln] = New_Vol(1, 1, this.Out_depth[ln], bias);



return this
}


func (this *layer) RegressionLayer(opt layer_def, ln int) *layer {
this.append_layer();
//var opt = opt || {};

// computed
this.num_inputs[ln] = opt.in_sx * opt.in_sy * opt.in_depth;
this.Out_depth[ln] = this.num_inputs[ln];
this.Out_sx[ln] = 1;
this.Out_sy[ln] = 1;
this.layer_type[ln] = "regression";

return  this

}


func (this *layer) SoftmaxLayer(opt layer_def, ln int) *layer {
this.append_layer();
//var opt = opt || {};

// computed
this.num_inputs[ln] = opt.in_sx * opt.in_sy * opt.in_depth;
this.Out_depth[ln] = this.num_inputs[ln];
this.Out_sx[ln] = 1;
this.Out_sy[ln] = 1;
this.layer_type[ln] = "softmax";


// not sure about this
return this


}

func (this *layer) SVMLayer(opt layer_def, ln int) *layer {
this.append_layer();

// computed
this.num_inputs[ln] = opt.in_sx * opt.in_sy * opt.in_depth;
this.Out_depth[ln] = this.num_inputs[ln];
this.Out_sx[ln] = 1;
this.Out_sy[ln] = 1;
this.layer_type[ln] = "svm";


// not sure about this
return this

}

func (this *layer) QuadTransformLayer(opt layer_def, ln int) *layer {

this.append_layer();


// computed
this.Out_sx[ln] = 1;
this.Out_sy[ln] = 1;

// linear terms, and then quadratic terms, of which there are 1/2*n*(n+1),
// (offdiagonals and the diagonal total) and arithmetic series.
// Actually never mind, lets not be fancy here yet and just include
// terms x_ix_j and x_jx_i twice. Half as efficient but much less
// headache.
this.Out_depth[ln] = opt.in_sx * opt.in_sy * opt.in_depth;
this.layer_type[ln] = "quadtransform";

return this
}

func (this *layer) FullyConnLayer(opt layer_def, ln int) *layer {
this.append_layer();
//var opt = opt || {};

// required
// ok fine we will allow 'filters' as the word as well
if opt.num_neurons != undefined_i {
this.Out_depth[ln] = opt.num_neurons


} else {
this.Out_depth[ln] = opt.filters
};

// optional
if opt.l1_decay_mul != undefined_f {

this.l1_decay_mul[ln] = opt.l1_decay_mul
} else {

this.l1_decay_mul[ln] = 0.0
}
if opt.l2_decay_mul != undefined_f {

this.l2_decay_mul[ln] = opt.l2_decay_mul
} else {
this.l2_decay_mul[ln] = 1.0;
}
// computed
this.num_inputs[ln] = opt.in_sx * opt.in_sy * opt.in_depth;



this.Out_sx[ln] = 1;
this.Out_sy[ln] = 1;
this.layer_type[ln] = "fc";
var bias float64
// initializations
if opt.bias_pref != undefined_f {

bias = opt.bias_pref
} else {
bias =  0.0
}


this.filters[ln] = make([]*Vol, 0);



//tv := New_Vol(1, 1, this.num_inputs[ln], undefined_f)

for i:=0;i<this.Out_depth[ln];i++ {


//temp := copy_temp_Vol(tv, undefined_f)


this.filters[ln] = append(this.filters[ln], New_Vol(1, 1, this.num_inputs[ln], undefined_f));


}
this.biases[ln] =  New_Vol(1, 1, this.Out_depth[ln], bias);

return this

}
func (this *layer) LocalResponseNormalizationLayer(opt layer_def, ln int) *layer {
//	var opt = opt || {};
this.append_layer();
// required
this.k[ln] = opt.k;
this.n[ln] = opt.n;
this.alpha[ln] = opt.alpha;
this.beta[ln] = opt.beta;

// computed
this.Out_sx[ln] = opt.in_sx;
this.Out_sy[ln] = opt.in_sy;
this.Out_depth[ln] = opt.in_depth;
this.layer_type[ln] = "lrn";

// checks
if(this.n[ln]%2 == 0) { println("WARNING n should be odd for LRN layer"); }

return this
}


func (this *layer) DropoutLayer(opt layer_def, ln int) *layer {
this.append_layer();
//var opt = opt || {};

// computed
this.Out_sx[ln] = opt.in_sx;
this.Out_sy[ln] = opt.in_sy;
this.Out_depth[ln] = opt.in_depth;
this.layer_type[ln] = "dropout";
if opt.drop_prob != undefined_f {
this.drop_prob[ln] = opt.drop_prob
} else {
t := 0.5
this.drop_prob[ln] = t
}

this.dropped[ln] = global.zeros(this.Out_sx[ln] * this.Out_sy[ln] * this.Out_depth[ln]);




return this
}


func (this *layer) append_layer() {

var inv *Vol;
var outv *Vol;
var bv *Vol;


this.sx = append(this.sx, undefined_i)
this.sy = append(this.sy, undefined_i)
this.stride = append(this.stride, undefined_i)
this.pad = append(this.pad, undefined_i)

this.n = append(this.n, undefined_i)
this.k = append(this.k, undefined_i)
this.alpha = append(this.alpha, undefined_f)
this.beta = append(this.beta, undefined_f)

this.in_sx = append(this.in_sx, undefined_i)
this.in_depth = append(this.in_depth, undefined_i)
this.in_sy = append(this.in_sy, undefined_i)
this.Out_depth = append(this.Out_depth, undefined_i)
this.Out_sx = append(this.Out_sx, undefined_i)
this.Out_sy = append(this.Out_sy, undefined_i)
this.layer_type = append(this.layer_type, undefined_s)
this.num_inputs = append(this.num_inputs, undefined_i)
this.in_act = append(this.in_act, inv)
this.out_act = append(this.out_act, outv)

this.drop_prob = append(this.drop_prob, dup_float((undefined_f)))
this.dropped = append(this.dropped, make([]float64, 0))
this.num_neurons = append(this.num_neurons, undefined_i)
this.filters = append(this.filters, make([]*Vol, 0))
this.l1_decay_mul = append(this.l1_decay_mul, undefined_f)
this.l2_decay_mul = append(this.l2_decay_mul, undefined_f)
this.bias_pref = append(this.bias_pref, undefined_f)
this.biases = append(this.biases, bv)
this.switches = append(this.switches, make([]float64, 0))
this.group_size = append(this.group_size, undefined_i)
this.es = append(this.es, make([]float64, (0.0)))

this.switchx = append(this.switchx, make([]float64, 0))
this.switchy = append(this.switchy, make([]float64, 0))



}


func (this *layer) InputLayer(opt layer_def, ln int) *layer {

this.append_layer();

//var opt = opt || {};

// this is a bit silly but lets allow people to specify either ins or outs
if opt.Out_sx != undefined_i {

this.Out_sx[ln] = opt.Out_sx
} else {
this.Out_sx[ln] =  opt.in_sx
};
if opt.Out_sy != undefined_i {
this.Out_sy[ln] = opt.Out_sy
} else {
this.Out_sy[ln] = opt.in_sy
};



if opt.Out_depth != undefined_i {

this.Out_depth[ln] = opt.Out_depth
}  else {
this.Out_depth[ln] = opt.in_depth

}

this.layer_type[ln] = "input";






// hmmmm.... is this right?
return this


}
func (this *layer) TanhLayer(opt layer_def, ln int) *layer {
//var opt = opt || {};
this.append_layer();
// computed
this.Out_sx[ln] = opt.in_sx;
this.Out_sy[ln] = opt.in_sy;
this.Out_depth[ln] = opt.in_depth;
this.layer_type[ln] = "tanh";

return  this


}

func (this *layer) SigmoidLayer(opt layer_def, ln int) *layer {


//var opt = opt || {};
this.append_layer();
// computed
this.Out_sx[ln] = opt.in_sx;
this.Out_sy[ln] = opt.in_sy;
this.Out_depth[ln] = opt.in_depth;
this.layer_type[ln] = "sigmoid";

return  this
}

func (this *layer) MaxoutLayer(opt layer_def, ln int) *layer {

this.append_layer();
// required
if  opt.group_size != undefined_i {

this.group_size[ln] =opt.group_size

} else {
this.group_size[ln] = 2
}


//var opt = opt || {};
this.append_layer();
// computed
this.Out_sx[ln] = opt.in_sx;
this.Out_sy[ln] = opt.in_sy;
this.Out_depth[ln] = int(math.Floor(float64(opt.in_depth) / float64( this.group_size[ln])))
this.layer_type[ln] = "maxout";

this.switches[ln] = global.zeros(this.Out_sx[ln]*this.Out_sy[ln]*this.Out_depth[ln]); // useful for backprop

return this
}

func (this *layer) ReluLayer(opt layer_def, ln int) *layer {

//var opt = opt || {};
this.append_layer();
// computed
this.Out_sx[ln] = opt.in_sx;
this.Out_sy[ln] = opt.in_sy;
this.Out_depth[ln] = opt.in_depth;
this.layer_type[ln] = "relu";


return this
}


func (this *layer) backward(y []float64, ln int, ys ystruct) float64 {







if this.layer_type[ln] ==  "svm" {



// compute and accumulate gradient wrt weights and bias of this layer
var x = this.in_act[ln];
x.Dw = global.zeros(len(x.w)); // zero out the gradient of input Vol

var yscore = x.w[int(y[0])]; // score of ground truth // hopefully y[0] is fine
var margin = 1.0;
var loss = 0.0;
for i:=0;i<this.Out_depth[ln];i++ {
if(-(yscore) + (x.w[i]) + margin > 0) {
// violating example, apply loss
// I love hinge loss, by the way. Truly.
// Seriously, compare this SVM code with Softmax forward AND backprop code above
// it's clear which one is superior, not only in code, simplicity
// and beauty, but also in practice.
x.Dw[i] = ((x.Dw[i]) + 1);
x.Dw[int(y[0])] = ((x.Dw[int(y[0])]) - 1);
loss += -(yscore) + (x.w[i]) + margin;
}
}

return loss;

}


if this.layer_type[ln] ==  "quadtransform" {



var V = this.in_act[ln];
V.Dw = global.zeros(len(V.w)); // zero out gradient wrt data
var V2 = this.out_act[ln];
var N = this.Out_depth[ln];
var Ni = V.depth;
for x:=0;x<V.sx;x++ {
for y:=0;y<V.sy;y++ {
for i:=0;i<N;i++ {
var chain_grad = V2.get_grad(x,y,i);
if(i<Ni) {
V.add_grad(x,y,i,(chain_grad));
} else {
var i0 = math.Floor(filter_float(float64(i-Ni))/filter_float(float64(Ni)));
var i1 = float64(i-Ni) - i0*float64(Ni);
V.add_grad(x,y,int(i0),(V.get(x,y,int(i1)))*(chain_grad));
V.add_grad(x,y,int(i1),(V.get(x,y,int(i0)))*(chain_grad));
}
}
}
}

}

if this.layer_type[ln] == "dropout" {


var V = this.in_act[ln]; // we need to set Dw of this
var chain_grad = this.out_act[ln];
var N = len(V.w);
V.Dw = global.zeros(N); // zero out gradient wrt data
//wg := sync.WaitGroup{}

//wg.Add(N)
for i:=0;i<N;i++ {
//go func(i int) {
//defer wg.Done()

if(this.dropped[ln][i] != 0.0) {
V.Dw[i] = chain_grad.Dw[i]; // copy over the gradient
}
//}(i)

}
//wg.Wait()
}


if this.layer_type[ln] == "maxout" {


var V = this.in_act[ln]; // we need to set Dw of this
var V2 = this.out_act[ln];
var N = this.Out_depth[ln];
V.Dw = global.zeros(len(V.w)); // zero out gradient wrt data

// pass the gradient through the appropriate switch
if(this.Out_sx[ln] == 1 && this.Out_sy[ln] == 1) {
for i:=0;i<N;i++ {
chain_grad := (V2.Dw[i]);
V.Dw[int((this.switches[ln][i]))] = (chain_grad);
}
} else {
// bleh okay, lets do this the hard way
var n=0; // counter for switches
for x:=0;x<V2.sx;x++ {
for y:=0;y<V2.sy;y++ {
for i:=0;i<N;i++ {
var chain_grad = V2.get_grad(x,y,i);
V.set_grad(x,y,int((this.switches[ln][n])),(chain_grad));
n++;
}
}
}
}

}

if this.layer_type[ln] == "pool" {



// pooling Layers have no parameters, so simply compute
// gradient wrt data here
var V = this.in_act[ln];
V.Dw = global.zeros(len(V.w)); // zero out gradient wrt data

//wg := sync.WaitGroup{}

//wg.Add(this.Out_depth[ln])
//var A = this.out_act[ln]; // computed in forward pass
//sem := thread_make(this.Out_depth[ln]) // threads on this layer might be a bad idea
n := 0;
for d:=0;d<this.Out_depth[ln];d++ {

//go func(d  int) {

//defer wg.Done()

var x = -this.pad[ln];
var y = -this.pad[ln];
x+=this.stride[ln]


for ax:=0; ax<this.Out_sx[ln]; ax++ {


if ax > 0 {
x+=this.stride[ln];// hopeflly this is right
}

y = -this.pad[ln];
y+=this.stride[ln]
for ay:=0; ay<this.Out_sy[ln]; ay++ {

if ay > 0 {
y+=this.stride[ln];// hopeflly this is right
}

chain_grad := this.out_act[ln].get_grad(ax,ay,d);



V.add_grad(int((this.switchx[ln][n])), int((this.switchy[ln][n])), d, (chain_grad));
n++;

}
}
//thread_signal(sem)
//}(d)


}
//wg.Wait()

//thread_wait(sem, this.Out_depth[ln])


}
if this.layer_type[ln] == "conv" {


// compute gradient wrt weights, biases and input data
this.in_act[ln].Dw = global.zeros(len(this.in_act[ln].w));
//V.Dw = global.zeros(len(V.w)); // zero out gradient wrt bottom data, we're about to fill it

//sem := thread_make(this.Out_depth[ln])

wg := sync.WaitGroup{}
//
wg.Add(this.Out_depth[ln])

//wg_wait_index := 0;
V  := *this.in_act[ln];

for d:=0;d<this.Out_depth[ln];d++ {
//go func(d int) {
//
//defer wg.Done()

f := *this.filters[ln][d];
var x = -this.pad[ln];
var y = -this.pad[ln];
x += this.stride[ln]

go func(y int, x int, d int, f Vol) {

defer wg.Done();
for ax := 0; ax < this.Out_sx[ln]; ax++ {

if ax > 0 {
x += this.stride[ln]; // hopeflly this is right
}
//sem := thread_make(this.Out_sy[ln])


y = -this.pad[ln];
y += this.stride[ln]
for ay := 0; ay < this.Out_sy[ln]; ay++ {

//go func(ay  int, y int, ax int, x int, d int) {

//defer wg.Done()

if ay > 0 {
y += this.stride[ln]; // hopeflly this is right
}


// conVolve and add up the gradients.
// could be more efficient, going for correctness first
var chain_grad = this.out_act[ln].get_grad(ax, ay, d); // gradient from above, from chain rule
for fx := 0; fx < f.sx; fx++ {
for fy := 0; fy < f.sy; fy++ {

for fd := 0; fd < f.depth; fd++ {
var oy = y + fy;
var ox = x + fx;
if (oy >= 0 && oy < V.sy && ox >= 0 && ox < V.sx) {
//

//
// forward prop calculated: a += f.get(fx, fy, fd) * V.get(ox, oy, fd);
//f.add_grad(fx, fy, fd, V.get(ox, oy, fd) * chain_grad);
//V.add_grad(ox, oy, fd, f.get(fx, fy, fd) * chain_grad);

// avoid function call overhead and use Vols directly for efficiency
var ix1 = ((V.sx * oy) + ox) * V.depth + fd;
var ix2 = ((f.sx * fy) + fx) * f.depth + fd;
f.Dw[ix2] = ((f.Dw[ix2]) + ((V.w[ix1]) * (chain_grad)));
V.Dw[ix1] = ((V.Dw[ix1]) + ((f.w[ix2]) * (chain_grad)));

}
}

}
}
this.biases[ln].Dw[d] = ((this.biases[ln].Dw[d]) + (chain_grad));

//thread_signal(sem)
//}(ay,y,ax,x,d)




}
//thread_wait(sem, this.Out_sy[ln])
}

this.filters[ln][d] = &f


}( y, x, d, f )
//}(d)

//thread_signal(sem)


}

wg.Wait();

this.in_act[ln] = &V
//
//thread_wait(sem, this.Out_depth[ln])

}


if this.layer_type[ln] == "input" {


}

if this.layer_type[ln] == "lrn" {

// evaluate gradient wrt data
var V = this.in_act[ln]; // we need to set Dw of this
V.Dw = global.zeros(len(V.w)); // zero out gradient wrt data
//var A = this.out_act; // computed in forward pass

var n2 = math.Floor(float64(this.n[ln])/2);

//wg := sync.WaitGroup{}

//wg.Add(V.sx)


for x:=0;x<V.sx;x++ {

//go func(x int) {
//	defer wg.Done()

for y := 0; y < V.sy; y++ {
for i := 0; i < V.depth; i++ {

var chain_grad = this.out_act[ln].get_grad(x, y, i);
var S = this.S_cache_[ln].get(x, y, i);
var SB = math.Pow((S), this.beta[ln]);
var SB2 = SB * SB;

// normalize in a window of size n
for j := math.Max(0, float64(i) - n2); j <= math.Min(float64(i) + n2, float64(V.depth) - 1); j++ {
var aj = V.get(x, y, int(j));
var g = -(aj) * this.beta[ln] * math.Pow((S), this.beta[ln] - 1) * this.alpha[ln] / filter_float(float64(this.n[ln])) * 2 * (aj);
if j == float64(i) {
g += SB
};
g /= SB2;
g *= (chain_grad);
V.add_grad(x, y, int(j), g);
}

}
}

//}(x)
}

//wg.Wait()
}

if this.layer_type[ln] == "fc" {

// this is really slow
var V = this.in_act[ln];
V.Dw = global.zeros(len(V.w)); // zero out the gradient in input Vol





//wg := sync.WaitGroup{}
sem := thread_make(this.Out_depth[ln])
// compute gradient wrt weights and data


//wg.Add(this.Out_depth[ln])

for i:=0;i<this.Out_depth[ln];i++ {


go func(i int) {
//defer wg.Done()

var tfi = this.filters[ln][i];
var chain_grad = this.out_act[ln].Dw[i];

for d:=0;d<this.num_inputs[ln];d++ {


V.Dw[d] = ((V.Dw[d]) + ((tfi.w[d])*(chain_grad))); // grad wrt input data
tfi.Dw[d] = ((tfi.Dw[d]) +  ((V.w[d]) * (chain_grad))); // grad wrt params
}


this.biases[ln].Dw[i] = ((this.biases[ln].Dw[i]) + (chain_grad)) ;
thread_signal(sem)
}(i)

}

//wg.Wait()
thread_wait(sem, this.Out_depth[ln])
}



if this.layer_type[ln] == "tanh" {

var V = this.in_act[ln]; // we need to set Dw of this
var V2 = this.out_act[ln];
var N = len(V.w);
V.Dw = global.zeros(N); // zero out gradient wrt data
////sem := thread_make(N)
//wg := sync.WaitGroup{}
//wg.Add(N)
for i:=0;i<N;i++ {

//	go func(i int) {
//	defer wg.Done()

var v2wi = (V2.w[i]);
V.Dw[i] = ((1.0 - v2wi * v2wi) * (V2.Dw[i]));

//	//thread_signal(sem)
//	}(i)

}

//wg.Wait()
////thread_wait(N, sem)
}

if this.layer_type[ln] == "sigmoid" {

var V = this.in_act[ln]; // we need to set Dw of this
var V2 = this.out_act[ln];
var N = len(V.w);
V.Dw = global.zeros(N); // zero out gradient wrt data

//wg := sync.WaitGroup{}
//wg.Add(N)

//sem := thread_make(N)
for  i:=0;i<N;i++ {
//go func(i int) {
//defer wg.Done()

var v2wi = (V2.w[i]);
V.Dw[i] = ((v2wi) * (1.0 - (v2wi)) * (V2.Dw[i]));
//		//thread_signal(sem)
//}(i)

}
//wg.Wait()

//this.in_act[ln] = &V
//this.out_act[ln] = &V2
////thread_wait(N, sem)

}

if this.layer_type[ln] == "relu" {

var V = this.in_act[ln]; // we need to set Dw of this
var V2 = this.out_act[ln];
var N = len(V.w);
V.Dw = global.zeros(N); // zero out gradient wrt data

////sem := thread_make(N)
//wg := sync.WaitGroup{}
//wg.Add(N)

for i:=0;i<N;i++ {
//	go func(i int) {
//	defer wg.Done()
//go func(i int) {

if (V2.w[i]) <= 0.0 {
V.Dw[i] = (0.0)
} else {
V.Dw[i] = V2.Dw[i]
};

//	//thread_signal(sem)
//}(i)


}
//wg.Wait()
////thread_wait(N, sem)
}


if this.layer_type[ln] == "softmax" {




// compute and accumulate gradient wrt weights and bias of this layer
var x = this.in_act[ln];


x.Dw = global.zeros(len(x.w)); // zero out the gradient of input Vol

wg := sync.WaitGroup{}
wg.Add(this.Out_depth[ln])

//sem := thread_make(this.Out_depth[ln])
for i:=0;i<this.Out_depth[ln];i++ {

//go func(i int) {

//defer wg.Done()


var indicator float64

if float64(i) == y[0] { // maybe change this to an i not a 0
indicator = 1.0
} else {
indicator = 0.0
}
var mul = -(indicator - (this.es[ln][i]));
x.Dw[i] = (mul);
//	//thread_signal(sem)
//}(i)


}

//wg.Wait()

//this.in_act[ln] = &x;

////thread_wait(this.Out_depth[ln], sem)

// loss is the class negative log likelihood
return -math.Log((this.es[ln][int(y[0])]));// maybe change this to an i not a 0 or something...
}


if this.layer_type[ln] == "regression" {



//sem := thread_make(this.Out_depth[ln])

// compute and accumulate gradient wrt weights and bias of this layer
var x = this.in_act[ln];
x.Dw = global.zeros(len(x.w)); // zero out the gradient of input Vol
var loss = 0.0;




if(ys.val == undefined_f && ys.dim == undefined_i) {
for i := 0; i < this.Out_depth[ln]; i++ {
////go func(i int) {
var dy = (x.w[i]) - y[i];
x.Dw[i] = (dy);
loss += 2 * dy * dy;
//thread_signal(sem)
//}(i)

}

//thread_wait(this.Out_depth[ln], sem)
//}
//else {
//	// assume it is a struct with entries .dim and .val
//	// and we pass gradient only along dimension dim to be equal to val
//	var i = y.dim;
//	var yi = y.val;
//	var dy = x.w[i] - yi;
//	x.Dw[i] = dy;
//	loss += 2*dy*dy;
//}


} else {



// assume it is a struct with entries .dim and .val
// and we pass gradient only along dimension dim to be equal to val
var i = ys.dim;
var yi = ys.val;
var dy float64 = (x.w[i]) - yi;
x.Dw[i] = (dy);
loss += 2*dy*dy;
//println(ys.dim)


}
return loss;


}




return 0.0
}


type ParamsAndGrads struct {

params []float64
grads []float64
l1_decay_mul float64
l2_decay_mul float64

}

func (this *layer) getParamsAndGrads(ln int) []ParamsAndGrads {


if this.layer_type[ln] == "svm" {

return make([]ParamsAndGrads, 0)
}


if this.layer_type[ln] == "dropout" {

return make([]ParamsAndGrads, 0)
}

if this.layer_type[ln] == "dropout" {
return make([]ParamsAndGrads, 0)
}
if this.layer_type[ln] == "maxout" {

return make([]ParamsAndGrads, 0)
}
if this.layer_type[ln] == "pool" {

return make([]ParamsAndGrads, 0)

}

if this.layer_type[ln] == "conv" {
var response []ParamsAndGrads = make([]ParamsAndGrads, 0);

for i:=0;i<this.Out_depth[ln];i++ {
response = append(response, ParamsAndGrads{params: this.filters[ln][i].w, grads: this.filters[ln][i].Dw, l2_decay_mul: this.l2_decay_mul[ln], l1_decay_mul: this.l1_decay_mul[ln]});
}
response = append(response, ParamsAndGrads{params: this.biases[ln].w, grads: this.biases[ln].Dw, l1_decay_mul: 0.0, l2_decay_mul: 0.0});
return response;

}


if this.layer_type[ln] == "input" {
return make([]ParamsAndGrads, 0)
}

if this.layer_type[ln] == "lrn" {

return make([]ParamsAndGrads, 0)

}


if this.layer_type[ln] == "fc" {
var response []ParamsAndGrads = make([]ParamsAndGrads, 0);
for i:=0;i<this.Out_depth[ln];i++ {
response = append(response, ParamsAndGrads{params: this.filters[ln][i].w, grads: this.filters[ln][i].Dw, l1_decay_mul: this.l1_decay_mul[ln], l2_decay_mul: this.l2_decay_mul[ln]});
}
response = append(response , ParamsAndGrads{params: this.biases[ln].w, grads: this.biases[ln].Dw, l1_decay_mul: (0.0) , l2_decay_mul: (0.0)});
return response;

}

if this.layer_type[ln] == "relu" {

return make([]ParamsAndGrads, 0)

}
if this.layer_type[ln] == "tanh" {
return make([]ParamsAndGrads, 0)
}
if this.layer_type[ln] == "sigmoid" {

return make([]ParamsAndGrads, 0)
}

if this.layer_type[ln] == "softmax" {

return make([]ParamsAndGrads, 0)
}

if this.layer_type[ln] == "regression" {

return make([]ParamsAndGrads, 0)
}

return make([]ParamsAndGrads, 0)
}

type jsonLayer struct {
Out_depth []int
Out_sx []int
Out_sy []int
layer_type []string

//

//sx []int
//sy  []int
//stride []int
//pad []int
//S_cache_ []*Vol
//in_sx []int
//in_depth []int
//in_sy []int
//Out_depth []int
//Out_sx []int
//Out_sy []int
//n []int
//k []int
//alpha []float64
//beta []float64
//layer_type []string
//num_inputs []int
//in_act []*Vol
//out_act []*Vol
//drop_prob []float64
//dropped [][]float64
//num_neurons []int
//filters [][]*Vol
//l1_decay_mul []float64
//l2_decay_mul []float64
//bias_pref []float64
//biases []*Vol
//switches []int
//group_size []int
//es [][]float64
//
}


func (this *Net) FromJSON(json_byte string)  {

this.fromJSON([]byte(json_byte));

}

func (this *LSTMNet) FromJSON(json_byte []byte)  {

/*
model map[string]*Vol
Hidden []*Vol
Cell []*Vol
needs_backprop bool
Backprop []func()
*/

this.Model = make(map[string]*Vol)

jsonc := make(map[string]interface{}, 0)

json.Unmarshal( json_byte, &jsonc)

if (jsonc["model"] != nil) {
model := jsonc["model"].(map[string]interface{})

for i,_ := range model {
var v = New_Vol(0,0,0,0);

v.fromJSON(model[i].(map[string]interface {}));

v.sx = 0
v.sy = 0
v.depth = 0

this.Model[i] = v
}
}


if (jsonc["Hidden"] != nil) {
Hidden := jsonc["Hidden"].([]interface{})

for i := 0; i < len(Hidden); i++ {
var v = New_Vol(0,0,0,0);


v.fromJSON(Hidden[i].(map[string]interface {}));


v.sx = 0
v.sy = 0
v.depth = 0

this.Hidden = append(this.Hidden, v)
}
}

if (jsonc["Cell"] != nil) {
Cell := jsonc["Cell"].([]interface{})
for i := 0; i < len(Cell); i++ {
var v = New_Vol(0,0,0,0);




v.fromJSON(Cell[i].(map[string]interface {}));

v.sx = 0
v.sy = 0
v.depth = 0

this.Cell = append(this.Cell, v)
}
}

this.needs_backprop = true
this.Backprop = make([]func(),0)

}

func (this *Net) fromJSON(json_byte []byte)  {

jsonc := make([]map[string]interface{}, 0)

json.Unmarshal( json_byte, &jsonc)



//this.Layers = make([]*layer, len(jsonc));






for i:=0;i<len(jsonc);i++ {


var Lj = jsonc[i]
var def layer_def = new_layer_def()

var t = Lj["layer_type"].(string);


if (t == "input") {
this.Layers = append(this.Layers,  this.layer.InputLayer(def, i));

}
if (t == "relu") {
this.Layers = append(this.Layers,  this.layer.ReluLayer(def, i));

}
if (t == "sigmoid") {
this.Layers = append(this.Layers,  this.layer.SigmoidLayer(def, i));

}
if (t == "tanh") {
this.Layers = append(this.Layers,  this.layer.TanhLayer(def, i));

}
if (t == "dropout") {
this.Layers = append(this.Layers,  this.layer.DropoutLayer(def, i));

}
if (t == "conv") {
this.Layers = append(this.Layers,  this.layer.ConvLayer(def, i));

}
if (t == "pool") {
this.Layers = append(this.Layers,  this.layer.PoolLayer(def, i));

}
if (t == "lrn") {
this.Layers = append(this.Layers,  this.layer.LocalResponseNormalizationLayer(def, i));

}
if (t == "softmax") {
this.Layers = append(this.Layers,  this.layer.SoftmaxLayer(def, i));

}
if (t == "regression") {
this.Layers = append(this.Layers,  this.layer.RegressionLayer(def, i));

}
if (t == "fc") {
this.Layers = append(this.Layers,  this.layer.FullyConnLayer(def, i));

}
if(t=="maxout") {
this.Layers = append(this.Layers, this.layer.MaxoutLayer(def, i));

}
if(t=="quadtransform") {
this.Layers = append(this.Layers,  this.layer.QuadTransformLayer(def, i));

}

}

//sem := thread_make(len(this.Layers))
wg := sync.WaitGroup{}

wg.Add(len(jsonc))

for i:=0;i<len(jsonc);i++ {

go func(i int) {

defer wg.Done()
var Lj = jsonc[i]
this.Layers[i].fromJSON(Lj, i);
//thread_signal(sem)
}(i)
}

wg.Wait()

//thread_wait(len(jsonc), sem)
}


func (this *Net) ToJSON() string {
	return string(this.toJSON());
}

func (this *Net) toJSON() []byte {

var Layers = make([]map[string]interface{}, len(this.Layers));
//sem := thread_make(len(this.Layers))

wg := sync.WaitGroup{}
wg.Add(len(this.Layers))

for i:=0;i<len(this.Layers);i++ {

go func(i int) {
defer wg.Done()
Layers[i] = (this.Layers[i].toJSON(i));

//thread_signal(sem)
}(i)

}

wg.Wait()
//thread_wait(len(this.Layers), sem)

j, _  := json.Marshal(Layers)

return j


}
func (this *Vol) toJSON() map[string]interface{} {

m := map[string]interface{}{"sx":this.sx,"sy":this.sy,"depth":this.depth,"w":this.w,"c":this.c,"d":this.d,"n":this.N}


return m
}

func (this *Vol) fromJSON(json map[string]interface{}) {


this.sx = int(json["sx"].(float64));
this.sy = int(json["sy"].(float64));


this.depth = int(json["depth"].(float64));

this.c = (json["c"].(float64));
this.d = int(json["d"].(float64));
this.N = int(json["n"].(float64));

if this.depth < 1 {
this.depth = 1
}
if this.sx < 1 {
this.sx = 1
}
if this.sy < 1 {
this.sy = 1
}

jw := json["w"].([]interface{})

var n = len(jw);

this.w = global.zeros(n);
this.Dw = global.zeros(n);
// copy over the elements.
for i:=0;i<n;i++  {

if jw[i] == nil {
this.w[i] = (0.0);
} else {
this.w[i] = (jw[i].(float64));
}
//this.w[i] = (jw[i].(float64));
}
}


func (this *layer) fromJSON(json map[string]interface{}, ln int) {




if json["layer_type"] == "svm" {

this.Out_depth[ln] = int(json["Out_depth"].(float64));
this.Out_sx[ln] = int(json["Out_sx"].(float64))
this.Out_sy[ln] = int(json["Out_sy"].(float64));
this.layer_type[ln] = json["layer_type"].(string);
this.num_inputs[ln] = int(json["num_inputs"].(float64));

}

if json["layer_type"] == "quadtransform" {

this.Out_depth[ln] = int(json["Out_depth"].(float64));
this.Out_sx[ln] = int(json["Out_sx"].(float64))
this.Out_sy[ln] = int(json["Out_sy"].(float64));
this.layer_type[ln] = json["layer_type"].(string);
}


if json["layer_type"] == "dropout" {
//test := this.DropoutLayer(new_layer_def(), ln);

this.Out_depth[ln] = int(json["Out_depth"].(float64));
this.Out_sx[ln] = int(json["Out_sx"].(float64))
this.Out_sy[ln] = int(json["Out_sy"].(float64));
this.layer_type[ln] = json["layer_type"].(string);
this.drop_prob[ln] = (json["drop_prob"].(float64));
this.dropped[ln] = global.zeros(this.Out_sx[ln] * this.Out_sy[ln] * this.Out_depth[ln]);
//println(len(this.dropped[ln]))
}


if json["layer_type"] == "maxout" {
this.Out_depth[ln] = int(json["Out_depth"].(float64));
this.Out_sx[ln] = int(json["Out_sx"].(float64))
this.Out_sy[ln] = int(json["Out_sy"].(float64));
this.layer_type[ln] = json["layer_type"].(string);
this.group_size[ln] = int(json["group_size"].(float64));
this.switches[ln] = global.zeros(this.group_size[ln]);
}

if json["layer_type"] == "input" {
this.Out_depth[ln] = int(json["Out_depth"].(float64));
this.Out_sx[ln] = int(json["Out_sx"].(float64))
this.Out_sy[ln] = int(json["Out_sy"].(float64));
this.layer_type[ln] = json["layer_type"].(string);

}



if json["layer_type"] == "fc" {
this.Out_depth[ln] = int(json["Out_depth"].(float64));
this.Out_sx[ln] = int(json["Out_sx"].(float64))
this.Out_sy[ln] = int(json["Out_sy"].(float64));
this.layer_type[ln] = json["layer_type"].(string);
this.num_inputs[ln] = int(json["num_inputs"].(float64));


if json["l1_decay_mul"].(float64) != undefined_f {
this.l1_decay_mul[ln] = json["l1_decay_mul"].(float64)
} else {
this.l1_decay_mul[ln] = 1.0
}
if json["l2_decay_mul"].(float64) != undefined_f {
this.l2_decay_mul[ln] = json["l2_decay_mul"].(float64)
} else {
this.l2_decay_mul[ln] = 1.0
}
jsonFilters := (json["filters"].([]interface{}))
this.filters[ln] = make([]*Vol, 0);
for i:=0;i<len(jsonFilters);i++ {
var v = New_Vol(0,0,0,0);
v.fromJSON(jsonFilters[i].(map[string]interface {}));
this.filters[ln] = append(this.filters[ln], v);
}




this.biases[ln] = New_Vol(0,0,0,0);
this.biases[ln].fromJSON(json["biases"].(map[string]interface{}));

}

if json["layer_type"] == "tanh" {
this.Out_depth[ln] = int(json["Out_depth"].(float64));
this.Out_sx[ln] = int(json["Out_sx"].(float64))
this.Out_sy[ln] = int(json["Out_sy"].(float64));
this.layer_type[ln] = json["layer_type"].(string);
}

if json["layer_type"] == "sigmoid" {
this.Out_depth[ln] = int(json["Out_depth"].(float64));
this.Out_sx[ln] = int(json["Out_sx"].(float64))
this.Out_sy[ln] = int(json["Out_sy"].(float64));
this.layer_type[ln] = json["layer_type"].(string);

}

if json["layer_type"] == "relu" {
this.Out_depth[ln] = int(json["Out_depth"].(float64));
this.Out_sx[ln] = int(json["Out_sx"].(float64))
this.Out_sy[ln] = int(json["Out_sy"].(float64));
this.layer_type[ln] = json["layer_type"].(string);

}
if json["layer_type"] == "softmax" {
this.Out_depth[ln] = int(json["Out_depth"].(float64));
this.Out_sx[ln] = int(json["Out_sx"].(float64))
this.Out_sy[ln] = int(json["Out_sy"].(float64));
this.layer_type[ln] = json["layer_type"].(string);
this.num_inputs[ln] = int(json["num_inputs"].(float64));

}

if json["layer_type"] == "regression" {
this.Out_depth[ln] = int(json["Out_depth"].(float64));
this.Out_sx[ln] = int(json["Out_sx"].(float64))
this.Out_sy[ln] = int(json["Out_sy"].(float64));
this.layer_type[ln] = json["layer_type"].(string);
this.num_inputs[ln] = int(json["num_inputs"].(float64));

}

if json["layer_type"] == "pool" {

this.Out_depth[ln] = int(json["Out_depth"].(float64));
this.Out_sx[ln] = int(json["Out_sx"].(float64))
this.Out_sy[ln] = int(json["Out_sy"].(float64));
this.layer_type[ln] = json["layer_type"].(string);
this.sx[ln] = int(json["sx"].(float64));
this.sy[ln] = int(json["sy"].(float64))
this.stride[ln] = int(json["stride"].(float64))
this.in_depth[ln] = int(json["in_depth"].(float64))
if int(json["pad"].(float64)) != undefined_i {
this.pad[ln] = int(json["pad"].(float64))
} else {
this.pad[ln] = 0
}
n := this.Out_sx[ln]*this.Out_sy[ln]*this.Out_depth[ln]
this.switchx[ln] = global.zeros(n); // need to re-init these appropriately
this.switchy[ln] = global.zeros(n);

}


if json["layer_type"] == "conv" {
this.Out_depth[ln] = int(json["Out_depth"].(float64));
this.Out_sx[ln] = int(json["Out_sx"].(float64))
this.Out_sy[ln] = int(json["Out_sy"].(float64));
this.layer_type[ln] = json["layer_type"].(string);
this.sx[ln] = int(json["sx"].(float64));
this.sy[ln] = int(json["sy"].(float64))
this.stride[ln] = int(json["stride"].(float64))
this.in_depth[ln] = int(json["in_depth"].(float64))
this.filters[ln] = make([]*Vol, 0)
if json["l1_decay_mul"] != nil {
this.l1_decay_mul[ln] = json["l1_decay_mul"].(float64)
} else {
this.l1_decay_mul[ln] = 0.0
}
if json["l2_decay_mul"] != nil {

this.l2_decay_mul[ln] = json["l2_decay_mul"].(float64)
}else {
this.l1_decay_mul[ln] = 0.0
}


if int(json["pad"].(float64)) != undefined_i {
this.pad[ln] = int(json["pad"].(float64))
} else {
this.pad[ln] = 0
}


var jsonFilters []interface{}
if json["filters"] != nil {

jsonFilters = (json["filters"].([]interface{}))
} else {

jsonFilters = make([]interface{}, 0)
}

this.filters[ln] = make([]*Vol, 0);
for i:=0;i<len(jsonFilters);i++ {


var v = New_Vol(0,0,0,0);
v.fromJSON(jsonFilters[i].(map[string]interface {}));
this.filters[ln] = append(this.filters[ln], v);

//fmt.Println(this.filters[ln][i].w) // here

}


this.biases[ln] = New_Vol(0,0,0,0);
if json["biases"] != nil {


this.biases[ln].fromJSON(json["biases"].(map[string]interface{}));
}


//this.in_act[ln] = this.in_act[ln-1];

}



}

func (this *LSTMNet) ToJSON() []byte {
m := map[string]interface{}{}


model  := make(map[string]interface{})

for i,_ := range this.Model {

json := make(map[string]interface{})

json = this.Model[i].toJSON()

model[i] = json//append(model,json);

}

m["model"] = model

	var Hidden []map[string]interface{}

for i:=0;i<len(this.Hidden);i++ {
Hidden = append(Hidden,this.Hidden[i].toJSON());
}

m["Hidden"] = Hidden


var Cell []map[string]interface{}

for i:=0;i<len(this.Cell);i++ {
Cell = append(Cell,this.Cell[i].toJSON());

}
m["Cell"] = Cell


/*
model map[string]*Vol
Hidden []*Vol
Cell []*Vol
needs_backprop bool
Backprop []func()
*/

j, _  := json.Marshal(m)

return j

}

func (this *layer) toJSON(ln int) map[string]interface{} {


//var json jsonLayer;



m := map[string]interface{}{}

if this.layer_type[ln] == "svm" {

m = map[string]interface{}{"Out_depth":this.Out_depth[ln], "Out_sx":this.Out_sx[ln], "Out_sy":this.Out_sy[ln], "layer_type":this.layer_type[ln],"num_inputs":this.num_inputs[ln] }

return m


}
if this.layer_type[ln] == "quadtransform" {
m = map[string]interface{}{"Out_depth":this.Out_depth[ln], "Out_sx":this.Out_sx[ln], "Out_sy":this.Out_sy[ln], "layer_type":this.layer_type[ln]}

return m

}
if this.layer_type[ln] == "dropout" {

m = map[string]interface{}{"Out_depth":this.Out_depth[ln], "Out_sx":this.Out_sx[ln], "Out_sy":this.Out_sy[ln], "layer_type":this.layer_type[ln],"drop_prob":this.drop_prob[ln]}

return m

}

if this.layer_type[ln] == "maxout" {

m = map[string]interface{}{"Out_depth":this.Out_depth[ln], "Out_sx":this.Out_sx[ln], "Out_sy":this.Out_sy[ln], "layer_type":this.layer_type[ln],"group_size":this.group_size[ln]}

return m
}


if this.layer_type[ln] == "input" {

m = map[string]interface{}{"Out_depth":this.Out_depth[ln], "Out_sx":this.Out_sx[ln], "Out_sy":this.Out_sy[ln], "layer_type":this.layer_type[ln]}

return m

}

if this.layer_type[ln] == "pool" {

m = map[string]interface{}{"sx":this.sx[ln],"sy":this.sy[ln],"stride":this.stride[ln],"Out_depth":this.Out_depth[ln],"in_depth":this.in_depth[ln], "Out_sx":this.Out_sx[ln], "Out_sy":this.Out_sy[ln], "layer_type":this.layer_type[ln],"pad":this.pad[ln] }

return m

}

//if this.layer_type[ln] == "pool" {
//
//	m = map[string]interface{}{"sx":this.sx[ln], "sy":this.sy[ln],"stride":this.stride[ln],"in_depth":this.in_depth[ln],"pad":this.pad[ln],"Out_depth":this.Out_depth[ln], "Out_sx":this.Out_sx[ln], "Out_sy":this.Out_sy[ln], "layer_type":this.layer_type[ln]}
//
//	return m
//
//
//}

if this.layer_type[ln] == "conv" {



var filters []map[string]interface{}

for i:=0;i<len(this.filters[ln]);i++ {
filters = append(filters,this.filters[ln][i].toJSON());

//fmt.Println(this.filters[ln][i].w)
//fmt.Println(filters[i]["w"])
}

biases := this.biases[ln].toJSON()

m = map[string]interface{}{"sx":this.sx[ln], "sy":this.sy[ln], "stride":this.stride[ln], "in_depth":this.in_depth[ln],"l1_decay_mul":this.l1_decay_mul[ln],"l2_decay_mul":this.l2_decay_mul[ln],"pad":this.pad[ln],"filters":filters, "biases":biases, "Out_depth":this.Out_depth[ln], "Out_sx":this.Out_sx[ln], "Out_sy":this.Out_sy[ln], "layer_type":this.layer_type[ln] }

return m

}


if this.layer_type[ln] == "softmax" {

m = map[string]interface{}{"Out_depth":this.Out_depth[ln], "Out_sx":this.Out_sx[ln], "Out_sy":this.Out_sy[ln], "layer_type":this.layer_type[ln],"num_inputs":this.num_inputs[ln] }

return m

}

if this.layer_type[ln] == "regression" {

m = map[string]interface{}{"Out_depth":this.Out_depth[ln], "Out_sx":this.Out_sx[ln], "Out_sy":this.Out_sy[ln], "layer_type":this.layer_type[ln],"num_inputs":this.num_inputs[ln] }

return m

}

if this.layer_type[ln] == "relu" {

m = map[string]interface{}{"Out_depth":this.Out_depth[ln], "Out_sx":this.Out_sx[ln], "Out_sy":this.Out_sy[ln], "layer_type":this.layer_type[ln]}

return m

}

if this.layer_type[ln] == "tanh" {

m = map[string]interface{}{"Out_depth":this.Out_depth[ln], "Out_sx":this.Out_sx[ln], "Out_sy":this.Out_sy[ln], "layer_type":this.layer_type[ln]}

return m
}


if this.layer_type[ln] == "sigmoid" {

m = map[string]interface{}{"Out_depth":this.Out_depth[ln], "Out_sx":this.Out_sx[ln], "Out_sy":this.Out_sy[ln], "layer_type":this.layer_type[ln]}

return m
}

if this.layer_type[ln] == "fc" {

var filters []map[string]interface{}

for i:=0;i<len(this.filters[ln]);i++ {
filters = append(filters,this.filters[ln][i].toJSON());
}


biases := this.biases[ln].toJSON()



m = map[string]interface{}{"biases":biases, "Out_depth":this.Out_depth[ln], "Out_sx":this.Out_sx[ln], "Out_sy":this.Out_sy[ln], "layer_type":this.layer_type[ln], "num_inputs":this.num_inputs[ln],"l1_decay_mul":this.l1_decay_mul[ln],"l2_decay_mul":this.l2_decay_mul[ln],"filters":filters }

return m


}

return m
}

func tanh (x float64) float64 {
var y = math.Exp(2 * x);
return (y - 1) / (y + 1);
}

type Prev struct {
	h []*Vol
	c []*Vol
	O *Vol
} 

func iToS (i int) string {
	return strconv.Itoa(i)
}

func sig(x float64) float64 {
    // helper function for computing sigmoid
    return 1.0/(1+math.Exp(-x));
  }

 type Solver struct {

	 decay_rate float64;
    smooth_eps float64;
    step_cache  map[string]*Vol;
 } 

 func NewSolver() *Solver {
	 S := new(Solver)
	 S.step_cache = make(map[string]*Vol)
	 S.smooth_eps = 0.01;
	 S.decay_rate = 0.01;
	 return S
 }

 func (this *Solver) Step(model *LSTMNet, step_size float64, regc float64, clipval float64) map[string]interface{} {
	 // perform parameter update
      var solver_stats = make(map[string]interface{});
      var num_clipped = 0;
      var num_tot = 0;
      for k,_ := range model.Model {
        if(model.Model[k] != nil ) {
          var m = model.Model[k]; // mat ref
          if !(this.step_cache[k] != nil) { 
			  mt := new(Vol) 
			  mt.Mat(m.N, m.d)
			  this.step_cache[k] = mt; 
			}
          var s = this.step_cache[k];
		  n:=len(m.w)

		//  wg := sync.WaitGroup{}
		//  wg.Add(n)
//
          for i:=0;i<n;i++ {
 	//go func(i int) {
			//	  defer wg.Done()
            // rmsprop adaptive learning rate
            var mDwi = m.Dw[i];
            s.w[i] = s.w[i] * this.decay_rate + (1.0 - this.decay_rate) * mDwi * mDwi;

            // gradient clip
            if(mDwi > clipval) {
              mDwi = clipval;
              num_clipped++;
            }
            if(mDwi < -clipval) {
              mDwi = -clipval;
              num_clipped++;
            }
            num_tot++;

            // update (and regularize)
            m.w[i] += - (step_size) * mDwi / math.Sqrt(s.w[i] + this.smooth_eps) - regc * m.w[i];
			//fmt.Println(m.w[i])
            m.Dw[i] = 0; // reset gradients for next iteration
		// }(i)
          }

		//  wg.Wait()
        }
      }
      solver_stats["ratio_clipped"] = num_clipped*1.0/num_tot;
      return solver_stats;

 }

func (this *LSTMNet) Softmax(m *Vol) *Vol {

       out := new(Vol) // probability volume
	   out.Mat(m.N, m.d);
       maxval := -999999.0;
	  n:=len(m.w)

	   wg := sync.WaitGroup{}
		  wg.Add(n)


      for i:=0;i<n;i++ { 
		    go func(i int) {
				  defer wg.Done()
		  if m.w[i] > maxval {
			  maxval = m.w[i]
			} 
			}(i)
		}

		wg.Wait()

      var s = 0.0;
	  n=len(m.w)
      for i:=0;i<n;i++ { 
        out.w[i] = math.Exp(m.w[i] - maxval);
        s += out.w[i];
      }

	   n=len(m.w)

	   wg = sync.WaitGroup{}
		  wg.Add(n)

      for i:=0;i<n;i++ { 
		  go func(i int) {
				  defer wg.Done()
		  out.w[i] /= s; 
		  }(i)
	}

		wg.Wait()

      // no backward pass here needed
      // since we will use the computed probabilities outside
      // to set gradients directly on m
      return out;
}


func (this *LSTMNet) Regression(m *Vol) *Vol {

		return m

    //    out := new(Vol) // probability volume
	//    out.Mat(m.N, m.d);
    //    maxval := -999999.0;
	//   n:=len(m.w)

	//    wg := sync.WaitGroup{}
	// 	  wg.Add(n)


    //   for i:=0;i<n;i++ { 
	// 	    go func(i int) {
	// 			  defer wg.Done()
	// 	  if m.w[i] > maxval {
	// 		  maxval = m.w[i] // 1<-
	// 		} 
	// 		}(i)
	// 	}

	// 	wg.Wait()

    //   var s = 0.0; // 2<-
	//   n=len(m.w)
    //   for i:=0;i<n;i++ { 
    //     out.w[i] = math.Exp(m.w[i] - maxval); // 3<-
    //     s += out.w[i];
    //   }

	//    n=len(m.w)

	//    wg = sync.WaitGroup{}
	// 	  wg.Add(n)

    //   for i:=0;i<n;i++ { 
	// 	  go func(i int) {
	// 			  defer wg.Done()
	// 	  out.w[i] /= s;  // 4<-
	// 	  }(i)
	// }

	// 	wg.Wait()

    //   // no backward pass here needed
    //   // since we will use the computed probabilities outside
    //   // to set gradients directly on m
    //   return out;
}

func (this *LSTMNet) Backward() {

for i:=len(this.Backprop)-1;i>=0;i-- {
        this.Backprop[i](); // tick!
      }
}


func (this *LSTMNet) tanh(m *Vol) *Vol {
	 // tanh nonlinearity
      var out = new(Vol)
	  out.Mat(m.N, m.d);
      var n = len(m.w);

	  wg := sync.WaitGroup{}
		  wg.Add(n)

      for i:=0;i<n;i++ { 
		  go func(i int) {
				  defer wg.Done()
        out.w[i] = math.Tanh(m.w[i]);
		  }(i)
      }
 		wg.Wait()


      if(this.needs_backprop) {
        var backward = func() {

			wg := sync.WaitGroup{}
		  wg.Add(n)
          for i:=0;i<n;i++ {

			  go func(i int) {
				  defer wg.Done()
           // grad for z = tanh(x) is (1 - z^2)
            var mwi = out.w[i];
            m.Dw[i] += (1.0 - mwi * mwi) * out.Dw[i];

			}(i)
          }
		  wg.Wait()

        }
        this.Backprop = append(this.Backprop, backward);
      }
      return out;
}

func(this *LSTMNet) eltmul(m1 *Vol, m2 *Vol) *Vol {
      assert(len(m1.w) == len(m2.w) , "etmul")

      var out = new(Vol)
	  out.Mat(m1.N, m1.d);
     n := len(m1.w)
	 wg := sync.WaitGroup{}
		  wg.Add(n)

	  for i:=0;i<n;i++ {
		   go func (i int)  {

			     defer wg.Done()
        out.w[i] = m1.w[i] * m2.w[i];
		 }(i)
      }

	  wg.Wait()
      if this.needs_backprop {
        var backward = func() {
          n:=len(m1.w)
	     wg := sync.WaitGroup{}
		 wg.Add(n)
		 
		  for i:=0;i<n;i++ {
			   go func (i int)  {
				   defer wg.Done()
            m1.Dw[i] += m2.w[i] * out.Dw[i];
            m2.Dw[i] += m1.w[i] * out.Dw[i];
				   }(i)
          }	

		  wg.Wait()
	
        }
        this.Backprop = append(this.Backprop, backward);
      }
      return out;
    }

func (this *LSTMNet) sigmoid(m *Vol) *Vol {
	  out := new(Vol)
	 out.Mat(m.N, m.d);
      n := len(m.w);

	   wg := sync.WaitGroup{}
		  wg.Add(n)

      for i:=0;i<n;i++ { 
		  go func (i int)  {
				   defer wg.Done()
        out.w[i] = sig(m.w[i]);
		  }(i)
}
wg.Wait()
  if(this.needs_backprop) {
        backward := func() {

			  wg := sync.WaitGroup{}
		 wg.Add(n)

          for i:=0;i<n;i++ {

			  	  go func (i int)  {
				   defer wg.Done()
            // grad for z = tanh(x) is (1 - z^2)
            var mwi = out.w[i];
            m.Dw[i] += mwi * (1.0 - mwi) * out.Dw[i];
					}(i)
          }
		  wg.Wait()
        }
        this.Backprop = append(this.Backprop, backward);
      }
      return out;
}

func (this *LSTMNet) add(m1 *Vol, m2 *Vol) *Vol {
 assert(len(m1.w) == len(m2.w), "add, m1 != m2")

	 out := new(Vol)
	 out.Mat(m1.N, m1.d);
	 	n := len(m1.w)

		  // wg := sync.WaitGroup{}
		 //// wg.Add(n)

      for i:=0;i<n;i++ {

		//    go func (i int)  {
			//	   defer wg.Done()
        out.w[i] = m1.w[i] + m2.w[i]
	//	}(i)
      }

	 //  wg.Wait()
      if this.needs_backprop {
         backward := func() {

			   wg := sync.WaitGroup{}
		  wg.Add(n)
			 n:=len(m1.w)
          for i:=0;i<n;i++ {

			  go func (i int)  {
				   defer wg.Done()
////
            m1.Dw[i] += out.Dw[i];
            m2.Dw[i] += out.Dw[i];

			}(i)
          }

		  wg.Wait()
        }
        this.Backprop = append(this.Backprop, backward);
      }
      return out;
}

func (this *LSTMNet) mul(m1 *Vol, m2 *Vol) *Vol {

	 // multiply matrices m1 * m2
      assert(m1.d == m2.N, `matmul dimensions misaligned`);

      var n = m1.N;
      var d = m2.d;
      var out = new(Vol)
	  out.Mat(n,d);
      for i:=0;i<m1.N;i++ { // loop over rows of m1
        for j:=0;j<m2.d;j++ { // loop over cols of m2
          var dot = 0.0;
          for k:=0;k<m1.d;k++ { // dot product loop
            dot += m1.w[m1.d*i+k] * m2.w[m2.d*k+j];
          }
          out.w[d*i+j] = dot;
        }
      }

      if(this.needs_backprop) {
       backward := func() {
          for i:=0;i<m1.N;i++ { // loop over rows of m1
            for j:=0;j<m2.d;j++ { // loop over cols of m2
              for k:=0;k<m1.d;k++ { // dot product loop
                 b := out.Dw[d*i+j];
                m1.Dw[m1.d*i+k] += m2.w[m2.d*k+j] * b;
                m2.Dw[m2.d*k+j] += m1.w[m1.d*i+k] * b;
              }
            }
          }
        }
        this.Backprop = append(this.Backprop, backward);
      }
      return out;
}

func (this *LSTMNet) ForwardLSTM (Hidden_sizes []int, x *Vol, prev Prev) Prev {
	  // forward prop for a single tick of LSTM
    // G is graph to append ops to
    // model contains LSTM parameters
    // x is 1D column vector with observation
    // prev is a struct containing Hidden and Cell
    // from previous iteration

	Hidden_prevs := make([]*Vol,0);
       Cell_prevs := make([]*Vol,0)

    if(prev.h == nil) {
       
      for d:=0;d<len(Hidden_sizes);d++ {
			
		R1 := new(Vol)
		R1.Mat(Hidden_sizes[d],1)
        Hidden_prevs = append(Hidden_prevs, R1); 

		R2 := new(Vol)
		R2.Mat(Hidden_sizes[d],1)
        Cell_prevs = append(Cell_prevs,  R2); 
	
      }

	
    } else {
       Hidden_prevs = prev.h;
       Cell_prevs = prev.c;
    }

     Hidden := make([]*Vol, 0);
     Cell := make([]*Vol, 0);

	  wg := sync.WaitGroup{}
		  wg.Add(len(Hidden_sizes))
		  
    for d:=0;d<len(Hidden_sizes);d++ {

			input_vector := x
		if d == 0 {
			input_vector = x
		} else {
			input_vector = Hidden[d-1];
		}
       Hidden_prev := Hidden_prevs[d];
       Cell_prev := Cell_prevs[d];


      // input gate
       h0 := this.mul(this.Model["Wix"+iToS(d)], input_vector);
       h1 := this.mul(this.Model["Wih"+iToS(d)], Hidden_prev);
      var input_gate = this.sigmoid(this.add(this.add(h0,h1),this.Model["bi"+iToS(d)]));

      // forget gate
      var h2 = this.mul(this.Model["Wfx"+iToS(d)], input_vector);
      var h3 = this.mul(this.Model["Wfh"+iToS(d)], Hidden_prev);
      var forget_gate = this.sigmoid(this.add(this.add(h2, h3),this.Model["bf"+iToS(d)]));

      // output gate
      var h4 = this.mul(this.Model["Wox"+iToS(d)], input_vector);
      var h5 = this.mul(this.Model["Woh"+iToS(d)], Hidden_prev);
      var output_gate = this.sigmoid(this.add(this.add(h4, h5),this.Model["bo"+iToS(d)]));

      // write operation on Cells
      var h6 = this.mul(this.Model["Wcx"+iToS(d)], input_vector);
      var h7 = this.mul(this.Model["Wch"+iToS(d)], Hidden_prev);
      var Cell_write = this.tanh(this.add(this.add(h6, h7),this.Model["bc"+iToS(d)]));

      // compute new Cell activation
      var retain_Cell = this.eltmul(forget_gate, Cell_prev); // what do we keep from Cell
      var write_Cell = this.eltmul(input_gate, Cell_write); // what do we write to Cell
      var Cell_d = this.add(retain_Cell, write_Cell); // new Cell contents

      // compute Hidden state as gated, saturated Cell activations
      var Hidden_d = this.eltmul(output_gate, this.tanh(Cell_d));

      Hidden = append(Hidden, Hidden_d);
      Cell = append(Cell, Cell_d);
    }

    // one decoder to outputs at end
    var output = this.add(this.mul(this.Model["Whd"], Hidden[len(Hidden) - 1]),this.Model["bd"]);

    // return Cell memory, Hidden representation and output
	p := Prev{}
	p.h = Hidden
	p.c = Cell
	p.O = output
	
    return p;
}


func (this *layer) forward(V *Vol, is_training bool, ln int) *Vol {




//source := rand.NewSource(time.Now().UnixNano())
//generator := rand.New(source)

if this.layer_type[ln] == "svm" {


this.in_act[ln] = V;
this.out_act[ln] = V; // nothing to do, output raw scores
return V;

}

if this.layer_type[ln] == "quadtransform" {


this.in_act[ln] = V;
var N = this.Out_depth[ln];
var Ni = V.depth;
var V2 = New_Vol(this.Out_sx[ln], this.Out_sy[ln], this.Out_depth[ln], 0.0);

//wg := sync.WaitGroup{}
//wg.Add(V.sx)

for x:=0;x<V.sx;x++ {
//	go func(x int) {

//	defer wg.Done()
for y := 0; y < V.sy; y++ {
for i := 0; i < N; i++ {
if (i < Ni) {
V2.set(x, y, i, (V.get(x, y, i))); // copy these over (linear terms)
} else {
var i0 float64 = math.Floor(filter_float(float64((i - Ni))) / filter_float(float64(Ni)));
var i1 float64 = float64(i - (Ni)) - i0 * float64(Ni);
V2.set(x, y, i, ((V.get(x, y, int(i0))) * (V.get(x, y, int(i1))))); // quadratic
}
}
}
//	}(x)
}
//wg.Wait()
this.out_act[ln] = V2;
return this.out_act[ln]; // dummy identity function for now


}
//fmt.Println(this.layer_type[ln])
if this.layer_type[ln] == "dropout" {

this.in_act[ln] = V;
//if(typeof(is_training)=='undefined') { is_training = false; } // default is prediction mode
var V2 = V.clone();
var N = len(V.w);
//println(N, len(this.dropped[ln]))
if(is_training) {
// do dropout
for i:=0;i<N;i++ {
if(rand.Float64()<(this.drop_prob[ln])) { V2.w[i]=(0.0); this.dropped[ln][i] = (1.0);  // this says false but I have set it to nil
}  else {this.dropped[ln][i] = (0.0);// this says false but I have set it to nil
}
}
} else {
// scale the activations during prediction
for i:=0;i<N;i++ { V2.w[i]= ((V2.w[i]) * (this.drop_prob[ln])); }
}
this.out_act[ln] = V2;
return this.out_act[ln]; // dummy identity function for now

}

if this.layer_type[ln] == "maxout" {


this.in_act[ln] = V;
var N = this.Out_depth[ln];
var V2 = New_Vol(this.Out_sx[ln], this.Out_sy[ln], this.Out_depth[ln], 0.0);

// optimization branch. If we're operating on 1D arrays we dont have
// to worry about keeping track of x,y,d coordinates inside
// input Volumes. In convnets we do :(
if(this.Out_sx[ln] == 1 && this.Out_sy[ln] == 1) {
///wg := sync.WaitGroup{}
//wg.Add(N)

for i:=0;i<N;i++ {
//go func(i int) {

//defer wg.Done()
var ix = i * this.group_size[ln]; // base index offset
var a = (V.w[ix]);
var ai = 0;
for j := 1; j < this.group_size[ln]; j++ {
var a2 = (V.w[ix + j]);
if (a2 > a) {
a = a2;
ai = j;
}
}
V2.w[i] = (a);
this.switches[ln][i] = (float64(ix + ai));

//}(i)
}
//wg.Wait()
} else {

//	wg := sync.WaitGroup{}
//	wg.Add(V.sx)
var n=0; // counter for switches
for x:=0;x<V.sx;x++ {
//	go func(x int) {

//defer wg.Done()
for y := 0; y < V.sy; y++ {
for i := 0; i < N; i++ {
var ix = i * this.group_size[ln];
var a = (V.get(x, y, ix));
var ai = 0;
for j := 1; j < this.group_size[ln]; j++ {
var a2 = (V.get(x, y, ix + j));
if (a2 > a) {
a = a2;
ai = j;
}
}
V2.set(x, y, i, a);
this.switches[ln][n] = (float64(ix + ai));
n++;
}
}
//	}(x)
}
//	wg.Wait()

}
this.out_act[ln] = V2;
return this.out_act[ln];

}


if this.layer_type[ln] == "pool" {


this.in_act[ln] = V;

var A = New_Vol(this.Out_sx[ln], this.Out_sy[ln], this.Out_depth[ln], 0.0);
//println(this.Out_sx[ln] ,"test")
var n=0; // a counter for switches

//wg := sync.WaitGroup{}
//wg.Add(this.Out_depth[ln])

for d:=0;d<this.Out_depth[ln];d++ {

//	go func(d int) {

//defer wg.Done()
var x = -this.pad[ln];
var y = -this.pad[ln];

x += this.stride[ln]; // hopeflly this is right
for ax := 0; ax < this.Out_sx[ln]; ax++ {

//println(x)

if ax > 0 {
x += this.stride[ln]; // hopeflly this is right
}

y = -this.pad[ln];

y += this.stride[ln]; // hopeflly this is right
for ay := 0; ay < this.Out_sy[ln]; ay++ {

if ay > 0 {
y += this.stride[ln]; // hopeflly this is right
}


//
//y+=this.stride[ln];
// conVolve centered at this particular location
a := (-99999.0); // hopefully small enough ;\
// a := undefined_f
winx := -1; winy := -1;


for fx := 0; fx < this.sx[ln]; fx++ {
for fy := 0; fy < this.sy[ln]; fy++ {
oy := y + fy;
ox := x + fx;


if (oy >= 0 && oy < V.sy && ox >= 0 && ox < V.sx) {

v := (V.get(ox, oy, d))
// perform max pooling and store pointers to where
// the max came from. This will speed up backprop
// and can help make nice visualizations in future


if ((v) > (a)) {
a = (v); winx = ox; winy = oy; }// else {println((v), (a))}
}
}
}

//	println(winx)
this.switchx[ln][n] = (float64(winx));
this.switchy[ln][n] = (float64(winy));
n++;

//println(a)
A.set(ax, ay, d, (a));
//y+=this.stride[ln];// hopeflly this is right
}
}
//}(d)
}

//wg.Wait()

this.out_act[ln] = A;
return this.out_act[ln];

}

if this.layer_type[ln] == "conv" {




this.in_act[ln] = V;

var A = New_Vol(this.Out_sx[ln], this.Out_sy[ln], this.Out_depth[ln], (0.0));

//wg := sync.WaitGroup{}
//wg.Add(this.Out_depth[ln])

for d:=0;d<this.Out_depth[ln];d++ {

//go func(d int) {

//defer wg.Done()

var f = this.filters[ln][d]; // this is the problem

var x = -this.pad[ln];
var y = -this.pad[ln];
x += this.stride[ln]
for ax := 0; ax < this.Out_sx[ln]; ax++ {

if ax > 0 {
x += this.stride[ln]
}

y = -this.pad[ln];
y += this.stride[ln]
for ay := 0; ay < this.Out_sy[ln]; ay++ {

if ay > 0 {
y += this.stride[ln]
}

// conVolve centered at this particular location
// could be bit more efficient, going for correctness first
var a = 0.0;
for fx := 0; fx < f.sx; fx++ {
for fy := 0; fy < f.sy; fy++ {
for fd := 0; fd < f.depth; fd++ {
var oy = y + fy; // coordinates in the original input array coordinates
var ox = x + fx;
if (oy >= 0 && oy < V.sy && ox >= 0 && ox < V.sx) {
//a += f.get(fx, fy, fd) * V.get(ox, oy, fd);
// avoid function call overhead for efficiency, compromise modularity :(

// the filters dont hav e a Vol

a += (f.w[int(
((f.sx * fy) + fx) *
(f.depth) +
(fd))]) *
(
V.w[int(((V.sx * oy) + ox) *
V.depth + ((fd)))]);
}
}
}
}
a += (this.biases[ln].w[d]);
A.set(ax, ay, d, (a));
}
}
//}(d)


}

//wg.Wait()
this.out_act[ln] = A;
return this.out_act[ln];
}

if this.layer_type[ln] == "input" {


this.in_act[ln] = V;
this.out_act[ln] = V;

return this.out_act[ln]; // dummy identity function for now
}


if this.layer_type[ln] == "lrn" {


this.in_act[ln] = V;

var A = V.cloneAndZero();
this.S_cache_[ln] = V.cloneAndZero();
var n2 = math.Floor(float64(this.n[ln])/2);

//wg := sync.WaitGroup{}
//wg.Add(V.sx)
for x:=0;x<V.sx;x++ {

//go func(x int) {

//	defer wg.Done()
for y := 0; y < V.sy; y++ {
for i := 0; i < V.depth; i++ {

var ai = V.get(x, y, i);

// normalize in a window of size n
var den = 0.0;
for j := math.Max(0, float64(i) - n2); j <= math.Min(float64(i) + n2, float64(V.depth) - 1); j++ {
var aa = V.get(x, y, int(j));
den += ((aa) * (aa))
}
den *= (this.alpha[ln]) / filter_float(float64(this.n[ln]));
den += float64(this.k[ln]);
this.S_cache_[ln].set(x, y, i, (den)); // will be useful for backprop
den = math.Pow(den, this.beta[ln]);
A.set(x, y, i, ((ai) / den));
}
}
//}(x)
}
//wg.Wait()

this.out_act[ln] = A;
return this.out_act[ln]; // dummy identity function for now

}

if this.layer_type[ln] == "dropout" {


this.in_act[ln] = V;
//if(is_training==false) { is_training = false; } // default is prediction mode
var V2 *Vol = new(Vol)

copyVol(V2, V)

//wg := sync.WaitGroup{}


var N = len(V.w);

//wg.Add(N)

if(is_training) {
// do dropout


for i:=0;i<N;i++ {

//	go func(i int) {
//
//		defer wg.Done()

t := rand.Float64()
t2 := (this.drop_prob[ln])
if ( t < t2) {
var t2 float64 = 0.0
V2.w[i] = t2;
var t float64 = 1.0
this.dropped[ln][i] = t; // this is meant to be true
} else {
var t float64 = 0.0
this.dropped[ln][i] = t // this is meant to be false
}

//	}(i)
}

//wg.Wait()
} else {


// scale the activations during prediction
for i:=0;i<N;i++ {
//	go func(i int) {

//		defer wg.Done()
//!
t1:= (V2.w[i])
t2:= (this.drop_prob[ln])
t3 := t1* t2
V2.w[i] = t3;
//}(i)
}
//	wg.Wait()
}
this.out_act[ln] = V2;
return this.out_act[ln]; // dummy identity function for now
}




if this.layer_type[ln] == "fc" {


	this.in_act[ln] = V;



var A =  New_Vol(1, 1, this.Out_depth[ln], 0.0);



var Vw []float64 = V.w;

//wg := sync.WaitGroup{}

sem := thread_make(this.Out_depth[ln])

for i:=0;i<this.Out_depth[ln];i++ {


//wg.Add(1)
go func(i int) {

//defer wg.Done()

var a float64= 0.0;
var wi []float64 = (this.filters[ln][i].w);



//println(this.num_inputs[ln] , "inputs")
for d:=0;d<this.num_inputs[ln];d++ {



a += (Vw[d]) * (wi[d]); // for efficiency use Vols directly for now


}

a += (this.biases[ln].w[i]);




A.w[i] = (a);

thread_signal(sem)
//fmt.Println((a), ln)
}(i)
}

//wg.Wait()
thread_wait(sem, this.Out_depth[ln])

this.out_act[ln] = A;

//println((this.out_act[ln].w[0]))
return this.out_act[ln];


}

if this.layer_type[ln] == "maxout" {

this.in_act[ln] = V;
N := this.Out_depth[ln];

t := 0.0
V2 := New_Vol(this.Out_sx[ln], this.Out_sy[ln], this.Out_depth[ln], t);
//wg := sync.WaitGroup{}
// optimization branch. If we're operating on 1D arrays we dont have
// to worry about keeping track of x,y,d coordinates inside
// input Volumes. In convnets we do :(
if this.Out_sx[ln] == 1 && this.Out_sy[ln] == 1 {

//wg.Add(N)
for  i:=0;i<N;i++ {

//go func(i int) {
//defer wg.Done()
var ix = i * this.group_size[ln]; // base index offset
var a = V.w[ix];
var ai = 0;
for j:=1;j<this.group_size[ln];j++ {
var a2 = V.w[ix+j];


if((a2) > (a)) {
a = a2;
ai = j;
}
}
V2.w[i] = a;
this.switches[ln][i] = (float64(ix + ai));

//	}(i)
}

//	wg.Wait()
} else {
var n=0; // counter for switches

//wg.Add(V.sx)
for x:=0;x<V.sx;x++ {

//	go func(x int) {
//	defer wg.Done()

for y:=0;y<V.sy;y++ {
for i:=0;i<N;i++ {

var ix = i * this.group_size[ln];
var a = V.get(x, y, ix);
temp_a := (a)
var ai = 0;
for j := 1; j < this.group_size[ln]; j++ {
var a2 = V.get(x, y, ix + j);


if ((a2) > temp_a) {
a = a2;
ai = j;
}
}

V2.set(x, y, i, (a));
this.switches[ln][n] = (float64(ix + ai));
n++;

}
}
//	}(x)
}

//wg.Wait()

}
this.out_act[ln] = V2;
return this.out_act[ln];

}


if this.layer_type[ln] == "relu" {

this.in_act[ln] = V;
var V2 *Vol = V.clone();

var N = len(V.w);

var V2w = V2.w;
//wg := sync.WaitGroup{}

//wg.Add(N)


for i:=0;i<N;i++ {

//go func(i int) {
//defer wg.Done()

if ((V2w[i]) < 0) {


V2w[i] = (0.0)
} else {

}; // threshold at 0
//}(i)
}

//wg.Wait()
this.out_act[ln] = V2;
return this.out_act[ln];
}
if this.layer_type[ln] == "tanh" {

this.in_act[ln] = V;
V2 := V.cloneAndZero();
N := len(V.w);
//wg := sync.WaitGroup{}
//wg.Add(N)
for i:=0;i<N;i++ {
//	go func(i int) {

//		defer wg.Done()
V2.w[i] = (tanh((V.w[i])))
//}(i)
}

//wg.Wait()
this.out_act[ln] = V2;
return this.out_act[ln];

}
if this.layer_type[ln] == "sigmoid" {

this.in_act[ln] = V;
var V2 = V.cloneAndZero();
var N = len(V.w);
var V2w []float64 = V2.w;
var Vw []float64  = V.w;
//wg := sync.WaitGroup{}

//wg.Add(N)

for  i:=0;i<N;i++ {

//		go func(i int) {

//	defer wg.Done()
V2w[i] = (1.0 / (1.0 + math.Exp(-(Vw[i]))))

//			}(i)
}

//	wg.Wait()
this.out_act[ln] = V2;
return this.out_act[ln];


}


if this.layer_type[ln] == "softmax" {

this.in_act[ln] = V;



var A = New_Vol(1, 1, this.Out_depth[ln], (0.0));


// compute max activation
var as []float64 = V.w;


var amax float64 = V.w[0];


for i := 1; i < this.Out_depth[ln]; i++ {

if ((as[i]) > (amax)) {
amax = as[i] // 1<-
};
}

// compute exponentials (carefully to not blow up)
var es = global.zeros(this.Out_depth[ln]);
var esum float64 = 0.0; // 2<-

for i := range es {
//for i := 0; i < this.Out_depth[ln]; i++ {

var e = math.Exp((as[i]) - (amax)); // 3<-
esum += e;
es[i] = (e);

}

// normalize and output to sum to one
for  i:=0;i<this.Out_depth[ln];i++ {

es[i] = ((es[i]) / esum);// 4<-
A.w[i] = es[i];

}


this.es[ln] = es; // save these for backprop
this.out_act[ln] = A;
return this.out_act[ln];
}


if this.layer_type[ln] == "regression" {

this.in_act[ln] = V;
this.out_act[ln] = V;
return V; // identity function

}


//var v *Vol
return V
}






//input layer etc....

//toJSON: function() {
//var json = {};
//json.Out_depth = this.Out_depth;
//json.Out_sx = this.Out_sx;
//json.Out_sy = this.Out_sy;
//json.layer_type = this.layer_type;
//return json;
//},
//fromJSON: function(json) {
//this.Out_depth = json.Out_depth;
//this.Out_sx = json.Out_sx;
//this.Out_sy = json.Out_sy;
//this.layer_type = json.layer_type;
//}


type Trainer struct  {



Net *Net
learning_rate float64
l1_decay float64
l2_decay float64
batch_size int;
method string // sgd/adagrad/adadelta/windowgrad

momentum float64
ro  float64
eps float64

k int; // iteration counter
gsum [][]float64; // last iteration gradients (used for momentum calculations)
xsum [][]float64; // used in adadelta
}


//type training_options struct  {
//
//	learning_rate float64
//	l1_decay float64
//	l2_decay float64
//	batch_size int
//	method string
//	momentum float64
//	ro float64
//	eps float64
//
//
//
//}
func  new_trainer(net *Net) *Trainer {

var this *Trainer = new(Trainer)

this.Net = net;

//var options = options || {};
this.learning_rate = 0.01
this.l1_decay = 0.0
this.l2_decay = 0.0
this.batch_size =  1
this.method = "sgd"// sgd/adagrad/adadelta/windowgrad

this.momentum = 0.9// used in adadelta
this.ro = 0.95// used in adadelta
this.eps = 1e-6

this.k = 0; // iteration counter
this.gsum = make([][]float64, 0); // last iteration gradients (used for momentum calculations)
this.xsum = make([][]float64, 0); // used in adadelta

return this

}

func  New_trainer_with_defs(net *Net, defs Trainer_def) Trainer {
	return new_trainer_with_defs(net, trainer_def(defs))
}

func  new_trainer_with_defs(net *Net, defs trainer_def) Trainer {

var this Trainer //= new(Trainer)





this.Net = net;

//var options = options || {};
if defs.learning_rate == undefined_f {
this.learning_rate = 0.01
}else {
this.learning_rate = defs.learning_rate
}
if defs.l1_decay == undefined_f {
this.l1_decay = 0.0
}else {
this.l1_decay = defs.l1_decay
}

if defs.l2_decay == undefined_f {
this.l2_decay = 0.0
} else {
this.l2_decay = defs.l2_decay
}
if defs.batch_size == undefined_i {
this.batch_size =  1
} else {
this.batch_size = defs.batch_size
}


if defs.method == undefined_s {
this.method = "sgd"// sgd/adagrad/adadelta/windowgrad
}else {
this.method = defs.method
}

if defs.momentum == undefined_f {
this.momentum = 0.9// used in adadelta
}else {
this.momentum = defs.momentum
}

if defs.eps == undefined_f {
this.eps = 1e-6
} else {
this.eps = defs.eps
}
if defs.ro == undefined_f {
this.ro = 0.95// used in adadelta
} else {
this.ro = defs.ro
}
//this.method = "sgd"// sgd/adagrad/adadelta/windowgrad



this.k = 0; // iteration counter
this.gsum = make([][]float64, 0); // last iteration gradients (used for momentum calculations)
this.xsum = make([][]float64, 0); // used in adadelta

return this

}

func (this Trainer) TrainNormalNet(x interface{}, y []float64) float64  {


	return this.train(x.(*Vol),y,undefined_ystruct)
}





func (this *Trainer) train(x *Vol, y []float64, ys ystruct) float64  {




//var x *Vol = duplicate_Vol(V)



//var start = new Date().getTime();
_ = this.Net.forward(x, true); // also set the flag that lets the net know we're just training





//var end = new Date().getTime();
//var fwd_time = end - start;

//var start = new Date().getTime();
cost_loss := this.Net.backward(y, ys);


var l2_decay_loss float64 = 0.0;
var l1_decay_loss float64 = 0.0;
//var end = new Date().getTime();
//var bwd_time = end - start;


this.k++;
if this.batch_size == 0 || this.k % this.batch_size == 0 {

var pglist []ParamsAndGrads = this.Net.getParamsAndGrads();



// initialize lists for accumulators. Will only be done once on first iteration
if len(this.gsum) == 0 && (this.method != "sgd" || this.momentum > 0.0) {

//	println((len(this.gsum)))



// only vanilla sgd doesnt need either lists
// momentum needs gsum
// adagrad needs gsum
// adadelta needs gsum and xsum





for i:=0;i<len(pglist);i++ {



this.gsum = append(this.gsum, global.zeros(len(pglist[i].params)));
if this.method == "adadelta" {
this.xsum = append(this.xsum, global.zeros(len(pglist[i].params)));
} else {
this.xsum = append(this.xsum, make([]float64, 0)); // conserve memory
}
}
}

//wg := sync.WaitGroup{}

//wg.Add(len(pglist))

// perform an update for all sets of weights
for i:=0;i<len(pglist);i++ {


//	go func(i int) {

//	defer wg.Done()

var pg = pglist[i]; // param, gradient, other options in future (custom learning rate etc)
var p = pg.params;
var g = pg.grads;


// learning rate for some parameters.
var l2_decay_mul float64
if pg.l2_decay_mul != undefined_f {
l2_decay_mul = pg.l2_decay_mul
} else {
l2_decay_mul = 1.0
}

var l1_decay_mul float64
if pg.l1_decay_mul != undefined_f {
l1_decay_mul = pg.l1_decay_mul
} else {
l1_decay_mul = 1.0
}

var l2_decay float64 = this.l2_decay * l2_decay_mul;
var l1_decay float64 = this.l1_decay * l1_decay_mul;

var plen int = len(p);


for j := 0; j < plen; j++ {

pj := (p[j])

l2_decay_loss += filter_float(l2_decay * (pj) * (pj)) / 2; // accumulate weight decay loss
l1_decay_loss += l1_decay * math.Abs((pj));


var l1grad float64
if (pj) > 0 {
l1grad = l1_decay * 1
} else {
l1grad = l1_decay * -1
}

var l2grad = l2_decay * (pj);

var gij = (l2grad + l1grad + (g[j])) / float64(this.batch_size); // raw batch gradient


if this.method == "adagrad" {
var gsumi = this.gsum[i];

// adagrad update
gsumi[j] = ((gsumi[j]) + gij * gij);
var dx = - this.learning_rate / math.Sqrt((gsumi[j]) + (this.eps)) * gij;
pj += ((dx));
} else if this.method == "windowgrad" {
var gsumi = this.gsum[i];

// this is adagrad but with a moving window weighted average
// so the gradient is not accumulated over the entire history of the run.
// it's also referred to as Idea #1 in Zeiler paper on Adadelta. Seems reasonable to me!
gsumi[j] = (this.ro * (gsumi[j]) + (1 - this.ro) * gij * gij);
var dx = - this.learning_rate / math.Sqrt((gsumi[j]) + this.eps) * gij; // eps added for better conditioning
pj += (dx);
} else if this.method == "adadelta" {
var gsumi = this.gsum[i];
var xsumi = this.xsum[i];
// assume adadelta if not sgd or adagrad
gsumi[j] = (this.ro * (gsumi[j]) + (1 - this.ro) * gij * gij);
var dx = - math.Sqrt(((xsumi[j]) + this.eps) / ((gsumi[j]) + this.eps)) * gij;
xsumi[j] = (this.ro * (xsumi[j]) + (1 - this.ro) * dx * dx); // yes, xsum lags behind gsum by 1.
pj += (dx);
} else {
// assume SGD
if (this.momentum > 0.0) {
var gsumi = this.gsum[i];

// momentum update
var dx = this.momentum * (gsumi[j]) - this.learning_rate * gij; // step
gsumi[j] = (dx); // back this up for next iteration of momentum
pj += (dx); // apply corrected gradient
} else {
// vanilla sgd
pj += - this.learning_rate * gij;
}
}
p[j] = (pj)
g[j] = (0.0); // zero out gradient so that we can begin accumulating anew
}


//}(i)

}

//wg.Wait();
}



//// appending softmax_loss for backwards compatibility, but from now on we will always use cost_loss
//// in future, TODO: have to completely redo the way loss is done around the network as currently
//// loss is a bit of a hack. Ideally, user should specify arbitrary number of loss functions on any layer
//// and it should all be computed correctly and automatically.
////return {fwd_time: fwd_time, bwd_time: bwd_time,
////l2_decay_loss: l2_decay_loss, l1_decay_loss: l1_decay_loss,
////cost_loss: cost_loss, softmax_loss: cost_loss,
////loss: cost_loss + l1_decay_loss + l2_decay_loss}


//V= x

//println(cost_loss)

this.Net.score = append(this.Net.score, cost_loss + l1_decay_loss + l2_decay_loss)

return cost_loss + l1_decay_loss + l2_decay_loss

}







type layer_def struct  {

sx int
sy int
stride int
pad int

n int
k int
alpha float64
beta float64
Type string
Out_sx int
Out_sy int
depth int
num_classes int
num_neurons int
bias_pref float64
activation string
tensor string
group_size int
drop_prob float64
in_sx int
in_depth int
in_sy int
Out_depth int
l1_decay_mul float64

l2_decay_mul float64
filters int
}

type Layer_def layer_def



func desugar(defs  []layer_def) []layer_def {

var new_defs []layer_def

for i:=0;i<len(defs);i++ {


def := defs[i];

if(def.Type=="softmax" || def.Type=="svm") {

// add an fc layer here, there is no reason the user should
// have to worry about this and we almost always want to
var df1 layer_def = new_layer_def()
df1.Type = "fc"
df1.num_neurons = def.num_classes

new_defs = append(new_defs, df1);



}


if(def.Type=="regression") {

// add an fc layer here, there is no reason the user should
// have to worry about this and we almost always want to

var df2 layer_def = new_layer_def()
df2.Type = "fc"
df2.num_neurons = def.num_neurons
new_defs = append(new_defs, df2);
}


if((def.Type=="fc" || def.Type=="conv") && def.bias_pref == undefined_f){ // maybe alter the 'nil'
def.bias_pref = 0.0;

if(def.activation != undefined_s && def.activation == "relu") {
def.bias_pref = 0.1; // relus like a bit of positive bias to get gradients early
// otherwise it's technically possible that a relu unit will never turn on (by chance)
// and will never get any gradient and never contribute any computation. Dead relu.
}
}

if(def.tensor != undefined_s) {
// apply quadratic transform so that the upcoming multiply will include
// quadratic terms, equivalent to doing a tensor product
if(def.tensor != undefined_s) {
var d3 layer_def = new_layer_def()
d3.Type  = "quadtransform"
new_defs = append(new_defs, d3);
}
}

new_defs = append(new_defs, def)

if(def.activation != undefined_s) {
if(def.activation=="relu") {
new_defs = append(new_defs,layer_def{Type:"relu"});
} else if (def.activation=="sigmoid") {
new_defs = append(new_defs, layer_def{Type:"sigmoid"});
} else if (def.activation=="tanh") {
new_defs = append(new_defs, layer_def{Type:"tanh"});
} else if (def.activation=="maxout") {
// create maxout activation, and pass along group size, if provided
gs := 2
if def.group_size == undefined_i {
gs = def.group_size
}
new_defs = append(new_defs, layer_def{Type:"maxout", group_size:gs});
} else { println("ERROR unsupported activation " + def.activation);
}
}


if(def.drop_prob != undefined_f && def.Type != "dropout") {

new_defs = append(new_defs, layer_def{Type:"dropout", drop_prob: def.drop_prob});
}



}


return new_defs;
}


func (this *Net) makeLayers (defs []layer_def) { // maybe no * on layer defs



// few checks for now
if(len(defs)<2) {println("ERROR! For now at least have input and softmax Layers.");}
if(defs[0].Type != "input") {println("ERROR! For now first layer should be input.");}



defs = desugar(defs);




for i := range defs {
//for i:=0;i<len(defs);i++ {

//	println(defs[i].Type);

var def layer_def = new_layer_def()




def = defs[i];


if(i>0) {

var prev *layer
prev = this.Layers[i-1];
def.in_sx = prev.Out_sx[i-1];
def.in_sy = prev.Out_sy[i-1];
def.in_depth = prev.Out_depth[i-1];

}





//println(def.Type)

switch(def.Type) {
case "fc": this.Layers = append(this.Layers,  this.layer.FullyConnLayer(def, i)); break;
case "lrn": this.Layers = append(this.Layers,  this.layer.LocalResponseNormalizationLayer(def,i)); break;
case "dropout": this.Layers = append(this.Layers,  this.layer.DropoutLayer(def, i)); break;
case "input": this.Layers = append(this.Layers,  this.layer.InputLayer(def, i)); break;
case "softmax": this.Layers = append(this.Layers,  this.layer.SoftmaxLayer(def, i)); break; //===
case "regression": this.Layers = append(this.Layers,  this.layer.RegressionLayer(def, i)); break;
case "conv": this.Layers = append(this.Layers,  this.layer.ConvLayer(def, i)); break;
case "pool": this.Layers= append(this.Layers,  this.layer.PoolLayer(def, i)); break;
case "relu": this.Layers = append(this.Layers,  this.layer.ReluLayer(def, i)); break;
case "sigmoid": this.Layers = append(this.Layers,  this.layer.SigmoidLayer(def, i)); break; //===
case "tanh": this.Layers = append(this.Layers,  this.layer.TanhLayer(def, i )); break;
case "maxout": this.Layers = append(this.Layers,  this.layer.MaxoutLayer(def, i)); break;
case "quadtransform": this.Layers  = append(this.Layers,  this.layer.QuadTransformLayer(def, i)); break;
case "svm": this.Layers  = append(this.Layers,  this.layer.SVMLayer(def, i)); break;
default: println("ERROR: UNRECOGNIZED LAYER TYPE!");
}
}

}


func MakeLayers (this *Net, defs []Layer_def)  {

	ld := []layer_def{}

	for _, v := range defs {
		ld = append(ld, layer_def(v))
	}


	this.makeLayers(ld)



}

func TainerToNet (trainer Trainer, net *Net) {
	trainer.Net = net
}

func New_trainer_def() Trainer_def {

	return Trainer_def(new_trainer_def())
}


func arrContains (arr []float64, elt float64) bool {
n:=len(arr)
for i:=0;i<n;i++ {
if i > 0 {
n=len(arr)
}
if(arr[i]==elt) {return true;}
}
return false;
}

func arrUnique (arr []float64) []float64 {
var b = []float64{};
n:=len(arr)
for i:=0;i<n;i++ {
if i > 0 {
n=len(arr)
}
if(!arrContains(b, arr[i])) {
b = append(b,arr[i]);
}
}
return b;
}


func (this *MagicNet) predict_soft(data *Vol) *Vol {



// forward prop the best networks
// and accumulate probabilities at last layer into a an output Vol
var nv = int(math.Min(float64(this.ensemble_size), float64(len(this.evaluated_candidates.self))));
if nv == 0 { return New_Vol(0,0,0, 0.0); } // not sure what to do here? we're not ready yet

//println(len(this.evaluated_candidates.self))

var xout *Vol
var n int;
for j:=0;j<nv;j++ {
var net = this.evaluated_candidates.self[j].net;

var x = net.forward(data, false);
//println((x.w[0]), "wtf")


if j==0 {
xout = x;
n = len(x.w);

} else {
// add it on
for d:=0;d<n;d++ {

xout.w[d] = (filter_float((xout.w[d])  + (x.w[d])));

}

}

}


// produce average
for d:=0;d<n;d++ {




xout.w[d] = (filter_float((xout.w[d]) / float64(n)));
//println((xout.w[d]))
}



return xout;
}

type maxmin_struct struct {
maxi int
mini int
maxv float64
minv float64
dv float64
}

func maxmin( w []float64 ) maxmin_struct {
var mm maxmin_struct
if len(w) == 0 { return mm; } // ... ;s
var maxv = (w[0]);
var minv = (w[0]);
var maxi = 0;
var mini = 0;
var n = len(w);
for i:=1;i<n;i++ {
gwi := (w[i])
if(gwi > maxv) { maxv = gwi; maxi = i; }
if(gwi < minv) { minv = gwi; mini = i; }
}
return maxmin_struct{maxi: maxi, maxv: maxv, mini: mini, minv: minv, dv:maxv-minv};
}

func (this *MagicNet) predict(data *Vol) int {
var predicted_label int

var xout = this.predict_soft(data);




if len(xout.w) != 0 {
var stats = maxmin(xout.w);

predicted_label = stats.maxi;
} else {
predicted_label = -1; // error out
}

//println(predicted_label)
return predicted_label;

}

func (this *MagicNet) evalValErrors() []float64 {

var vals []float64 = []float64{};
var fold = this.folds[this.foldix]; // activ // e fold


for k:=0;k<len(this.candidates.self);k++ {
var net *Net = this.candidates.self[k].net;
var v float64 = 0.0;
for q:=0;q<len(fold.test_ix);q++ {
var x = this.data[fold.test_ix[q]];
var l = this.labels[fold.test_ix[q]];


//t1 := (x.w[0])
//println((x.w[0]))



//println(t1 == (x.w[0]))
//println(get_
// float(x.w[0]))

//	println(len(x.w))
net.forward(x, false)
var yhat = net.getPrediction();

//println(yhat)


if float64(yhat) == l {

v += 1.0
} else {

v += 0.0
}
//v += (yhat === l ? 1.0 : 0.0); // 0 1 loss
}



v /= filter_float(float64(len(fold.test_ix))); // normalize


v = (v)




vals = append(vals, v);
}


return vals;

}

func (s *candidates) Len() int {

return len(s.self)

}

func (s *candidates) Swap(i, j int) {
s.self[i], s.self[j] = s.self[j], s.self[i]

}




func (s *candidates) Less(a, b int) bool {


if ((s.self[a].accv / float64(len(s.self[a].acc))) > (s.self[b].accv / float64(len(s.self[b].acc)))) {

//println(a,b,"down")
//println((s.self[a].accv / float64(len(s.self[a].acc))), (s.self[b].accv / float64(len(s.self[b].acc))), a, b)
return true
} else {


return false
}

}

//func cand_sort (a *candidates, b  []*candidate) bool {
//	// i hope this works
//	if ((a.accv / float64(len(a.acc))) > (b.accv / float64(len(b.acc)))) {
//		return false
//	} else {
//		return true
//	}
//	//? -1 : 1;
//}
func (this *MagicNet) toJSON() []byte {
var nv = int(math.Min(float64(this.ensemble_size), float64(len(this.evaluated_candidates.self))));
//
json_net := make([]string, 0)
for i:=0;i<nv;i++ {
json_net = append(json_net,string(this.evaluated_candidates.self[i].net.toJSON()));
}
json_s := map[string]interface{}{};
json_s["nets"] = json_net

j, _  := json.Marshal(json_s)

return j;

}


func (this *MagicNet) fromJSON(json_byte []byte) {

json_s := make([]map[string]interface{}, 0)

json.Unmarshal(json_byte, &json_s)


this.ensemble_size = len(json_s);
this.evaluated_candidates = new(candidates)
for i:=0;i<this.ensemble_size;i++ {
var net = new(Net);
net.fromJSON([]byte(json_s[i]["nets"].(string)));
var dummy_candidate candidate
dummy_candidate.net = net;
this.evaluated_candidates.self = append(this.evaluated_candidates.self, &dummy_candidate);
}



}

func (this *MagicNet) step() {



// run an example through current candidate
this.iter++;

// step all candidates on a random data point


var fold = this.folds[this.foldix]; // active fold




//fold.train_ix = make([]int, len(fold.train_ix))


r:= randi(0, len(fold.train_ix))

//println(r, len(fold.train_ix))
var dataix = fold.train_ix[r];

////sem := thread_make(len(this.candidates.self))

for k:=0;k<len(this.candidates.self);k++ {

////go func(k int) {
var x = this.data[dataix];
var l = this.labels[dataix];



this.candidates.self[k].trainer.train(x, []float64{l}, ystruct{dim:undefined_i,val:undefined_f});

//	//thread_signal(sem)

//}(k)
}

////thread_wait(len(this.candidates.self), sem)

// process consequences: sample new folds, or candidates
var lastiter int = (this.num_epochs * len(fold.train_ix));


if(this.iter >= lastiter) {
// finished evaluation of this fold. Get final validation
// accuracies, record them, and go on to next fold.



var val_acc = this.evalValErrors();
for k:=0;k<len(this.candidates.self);k++ {
var c = this.candidates.self[k];
c.acc = append(c.acc, (val_acc[k]));

c.accv += val_acc[k];
}
this.iter = 0; // reset step number
this.foldix++; // increment fold


// maybe add this back
//if(this.finish_fold_callback !== null) {
//this.finish_fold_callback();
//}

if(this.foldix >= len(this.folds)) {
// we finished all folds as well! Record these candidates
// and sample new ones to evaluate.

//sem := thread_make(len(this.candidates.self))








for k:=0;k<len(this.candidates.self);k++ {

//go func(k int) {


this.evaluated_candidates.self = append(this.evaluated_candidates.self,  (this.candidates.self[k]))


//thread_signal(sem)
//}(k)

}


//thread_wait(len(this.candidates.self), sem)
// sort evaluated candidates according to accuracy achieved

//println(this.evaluated_candidates.self)

//println(this.evaluated_candidates.self[0].accv)


sort.Sort(this.evaluated_candidates); // this seems to be working.



//println(this.evaluated_candidates.self[0].accv)
//println(this.evaluated_candidates.self)

// and clip only to the top few ones (lets place limit at 3*ensemble_size)
// otherwise there are concerns with keeping these all in memory
// if MagicNet is being evaluated for a very long time
if(len(this.evaluated_candidates.self) > 3 * this.ensemble_size) {
this.evaluated_candidates.self = this.evaluated_candidates.self[0: 3 * this.ensemble_size];
}
// maybe add this back
//if(this.finish_batch_callback !== null) {
//this.finish_batch_callback();
//}



//	score := 0
//
//	sem = thread_make(len(this.labels))
//
//
//	for g := 0; g < len(this.labels); g ++{
//
//		//go func(g int) {
//		//println(this.predict(this.data[g]))
//		//println((this.data[g].w[0]), (this.data[g].w[1]))
//		if (float64((this.predict_soft(this.data[g]).w[int(this.labels[g])]))> 0.5){
//
//
//			score ++
//		}
//		//thread_signal(sem)
//
//	}(g)
//}
//
//	//thread_wait(len(this.labels), sem)
//
//	println(score, len(this.labels))


this.sampleCandidates(); // begin with new candidates
this.foldix = 0; // reset this


} else {
// we will go on to another fold. reset all candidates nets
for k:=0;k<len(this.candidates.self);k++ {
var c = this.candidates.self[k];
var net = new(Net);
net.makeLayers(c.layer_defs);
var trainer = new_trainer_with_defs(net, c.trainer_def);
c.net = net;
c.trainer = &trainer;
}
}



}


}

type magic_net_opts struct {
train_ratio float64
num_folds int
num_candidates int
num_epochs int
ensemble_size int
batch_size_min int
batch_size_max int
l2_decay_min float64
l2_decay_max float64
learning_rate_min float64
learning_rate_max float64
momentum_min float64
momentum_max float64
neurons_min int
neurons_max int
}

func new_magic_net_opts() magic_net_opts{

var this magic_net_opts
this.train_ratio = 0.7

this.num_folds = 10
this.num_candidates = 50
this.num_epochs = 50

this.ensemble_size = 10
this.batch_size_min = 10
this.batch_size_max = 300

this.l2_decay_max = 0.0
this.l2_decay_min = -4.0

this.learning_rate_min = -4.0
this.learning_rate_max = 0.0

this.momentum_max = 0.9
this.momentum_min = 0.9

this.neurons_min = 5
this.neurons_max = 30

return this

}

func (this *MagicNet) setup (data []*Vol, labels []float64, opts magic_net_opts) *MagicNet {

//var this *MagicNet = new(MagicNet)


this.data = data
this.labels = labels


this.train_ratio = opts.train_ratio

this.num_folds = opts.num_folds
this.num_candidates = opts.num_candidates
this.num_epochs = opts.num_epochs

this.ensemble_size = opts.ensemble_size
this.batch_size_min = opts.batch_size_min
this.batch_size_max = opts.batch_size_max

this.l2_decay_max = opts.l2_decay_max
this.l2_decay_min = opts.l2_decay_min

this.learning_rate_min = opts.learning_rate_min
this.learning_rate_max = opts.learning_rate_max

this.momentum_max = opts.momentum_max
this.momentum_min = opts.momentum_min

this.neurons_min = opts.neurons_min
this.neurons_max = opts.neurons_max



this.folds = make([]*fold, 0); // flush folds, if any
//this.folds = append(this.folds ,&fold{train_ix: []int{undefined_i}, test_ix: []int{undefined_i}});


this.candidates = new(candidates)
this.evaluated_candidates = new(candidates)
this.unique_labels = arrUnique(labels);
this.iter = 0
this.foldix = 0



if(len(this.data) > 0) {

this.sampleFolds();
this.sampleCandidates();
}

return this
}



func randperm(n int) []int {

//source := rand.NewSource(time.Now().UnixNano())
//generator := rand.New(source)


return rand.Perm(n)
//
//	//
//	var i int = n
//	var j float64= 0.0
//	var temp float64;
//	var array []float64 = make([]float64, n);
//for q:=0;q<n;q++ {array[q]=q};
//for i = i; i >= 0; i-- {
//j = math.Floor(rand.Float64() * (i+1));
//temp = array[i];
//array[i] = array[j];
//array[j] = temp;
//}
//return array;
}

func (this *MagicNet) sampleCandidates() {

this.candidates = new(candidates); // flush, if any
for i:=0;i<this.num_candidates;i++ {
var cand = this.sampleCandidate();
this.candidates.self = append(this.candidates.self, cand);
}
}


func (this *MagicNet) sampleFolds() {

var N = len(this.data);
var num_train = int(math.Floor(this.train_ratio * float64(N)));



this.folds = make([]*fold, 0); // flush folds, if any
for i:=0;i<this.num_folds;i++ {
var p = randperm(N);


this.folds = append(this.folds ,&fold{train_ix: p[0:num_train], test_ix: p[num_train:N]});
//println(this.folds[len(this.folds)-1].test_ix[0])
}




}



func new_cand(layer_defs []layer_def, trainer_def trainer_def, trainer Trainer, net *Net) candidate {
var cand candidate //= new(cand)
cand.acc = make([]float64, 0)
cand.accv = 0; // this will maintained as sum(acc) for convenience
cand.layer_defs = layer_defs;
cand.trainer_def = trainer_def;
cand.net = net;
cand.trainer = &trainer;
return cand

}
func randf (a float64, b float64) float64 {


//source := rand.NewSource(time.Now().UnixNano())
//generator := rand.New(source)

return rand.Float64()*(b-a)+a; }
func randi(a int, b int) int {

//source := rand.NewSource(time.Now().UnixNano())
//generator := rand.New(source)
return int(math.Floor(rand.Float64()*float64((b-a)+a))); }
func randn(mu float64, std float64) float64 { return mu+global.gaussRandom()*std; }

func weightedSample(lst []float64, probs []float64) float64 {
var p = randf(0, 1.0);
var cumprob = 0.0;
n := len(lst)
for k:=0;k<n;k++ {
if k > 0 {
n=len(lst)
}

cumprob += probs[k];
if(p < cumprob) { return lst[k]; }
}

return undefined_f

}

func new_trainer_def() trainer_def {
var this trainer_def //= new(trainer_def)
this.method = undefined_s
this.batch_size = undefined_i
this.eps = undefined_f
this.l1_decay = undefined_f
this.l2_decay = undefined_f
this.method = undefined_s
this.learning_rate = undefined_f
this.momentum = undefined_f
this.ro = undefined_f

return this
}

func (this *MagicNet) sampleCandidate() *candidate {

//source := rand.NewSource(time.Now().UnixNano())
//generator := rand.New(source)

var input_depth int = len(this.data[0].w);
var num_classes int = len(this.unique_labels);

// sample network topology and hyperparameters
var layer_defs []layer_def = make([]layer_def, 0)
//new_layer_def();
layer_defs = append(layer_defs, new_layer_def());
layer_defs[len(layer_defs)-1].Type = "input"
layer_defs[len(layer_defs)-1].Out_sx = 1
layer_defs[len(layer_defs)-1].Out_sy = 1
layer_defs[len(layer_defs)-1].Out_depth = input_depth

//	layer_defs{type:"input", Out_sx:1, Out_sy:1, Out_depth: input_depth}
var nl float64 = weightedSample([]float64{0,1,2,3}, []float64{0.2, 0.3, 0.3, 0.2}); // prefer nets with 1,2 Hidden Layers
for q :=0;q<int(nl);q++ {
var ni = randi(this.neurons_min, this.neurons_max);
var act string = []string{"tanh","maxout", "relu"}[randi(0,3)];
//var act string = []string{"tanh","maxout","relu"}[randi(0,3)];
if(randf(0,1)<0.5) {
var dp = rand.Float64();
layer_defs = append(layer_defs, new_layer_def());
layer_defs[len(layer_defs)-1].Type = "fc"
layer_defs[len(layer_defs)-1].num_neurons = ni
layer_defs[len(layer_defs)-1].activation = act
layer_defs[len(layer_defs)-1].drop_prob = dp
//layer_defs = append(layer_defs, {type:"fc", num_neurons: ni, activation: act, drop_prob: dp});
} else {
layer_defs = append(layer_defs, new_layer_def());
layer_defs[len(layer_defs)-1].Type = "fc"
layer_defs[len(layer_defs)-1].num_neurons = ni
layer_defs[len(layer_defs)-1].activation = act

//layer_defs.push({type:"fc", num_neurons: ni, activation: act});
}
}
layer_defs = append(layer_defs, new_layer_def());
layer_defs[len(layer_defs)-1].Type = "softmax"
layer_defs[len(layer_defs)-1].num_classes = num_classes
//layer_defs.push({type:"softmax", num_classes: num_classes});
var net = new(Net);

net.makeLayers(layer_defs);

// sample training hyperparameters
var bs = randi(this.batch_size_min, this.batch_size_max); // batch size
var l2 = math.Pow(10, randf(this.l2_decay_min, this.l2_decay_max)); // l2 weight decay
var lr = math.Pow(10, randf(this.learning_rate_min, this.learning_rate_max)); // learning rate
var mom = randf(this.momentum_min, this.momentum_max); // momentum. Lets just use 0.9, works okay usually ;p
var tp = randf(0,1); // trainer type
var trainer_def trainer_def = new_trainer_def();
if(tp<0.33) {
trainer_def.method = "adadelta"
trainer_def.batch_size = bs
trainer_def.l2_decay = l2
//trainer_def = {method:"adadelta", batch_size:bs, l2_decay:l2};

} else if(tp<0.66) {
trainer_def.method = "adagrad"
trainer_def.learning_rate = lr
trainer_def.batch_size = bs
trainer_def.l2_decay = l2
//trainer_def = {method:"adagrad", learning_rate: lr, batch_size:bs, l2_decay:l2};
} else {

trainer_def.method = "sgd"
trainer_def.learning_rate = lr
trainer_def.momentum = mom
trainer_def.l2_decay = l2


//trainer_def = {method:"sgd", learning_rate: lr, momentum: mom, batch_size:bs, l2_decay:l2};
}

var trainer Trainer = new_trainer_with_defs(net, trainer_def);

var cand candidate = new_cand(layer_defs, trainer_def, trainer, net)
//cand.acc = [];
//cand.accv = 0; // this will maintained as sum(acc) for convenience
//cand.layer_defs = layer_defs;
//cand.trainer_def = trainer_def;
//cand.net = net;
//cand.trainer = trainer;
return &cand;


}



func thread_make(N int) chan empty {
return make(chan empty)
}

func thread_signal(sem chan empty)  {
sem <- empty{}
}

func thread_wait(sem chan empty, N int) {

for i := 0; i < N; i++ { <-sem }
}



type cnnutil struct {

window *win

}

func  new_cnnutil() *cnnutil {

var this *cnnutil = new(cnnutil)
this.window = new(win)
return this
}

type win struct {

v []float64
size int
minsize int
sum float64

}




func new_window(size int, minsize int) *win {
var this *win = new(win)
this.v = make([]float64,0 );
if size == undefined_i {
this.size  = 100
} else {
this.size = size
}
if (minsize) == undefined_i {
this.minsize = 20
} else {
this.minsize = minsize
}
this.sum = 0;


return this
}

func (this *win) add(x float64) {
this.v = append(this.v, (x));
this.sum += x;
if(len(this.v)>this.size) {
var xold float64
xold, this.v = (this.v[0]), this.v[1:] //shift()
//var xold = this.v.shift();
this.sum -= xold;
}
}

func (this *win) get_average() float64 {
if(len(this.v) < this.minsize) {return -1
} else { return this.sum/float64(len(this.v))}
}

func (this *win) reset() {
this.v = make([]float64, 0);
this.sum = 0;
}

func ft2 (x float64, d int) string {

if d == undefined_i { d = 5 }
var dd float64 = 1.0 * math.Pow(10, float64(d));
return  strconv.FormatFloat(math.Floor(x*dd)/dd, 'E', -1, 64);

}




type brain_opt struct {

temporal_window int
experience_size int
start_learn_threshold float64
gamma float64
learning_steps_total int
learning_steps_burnin int
epsilon_min float64
epsilon_test_time float64
random_action_distribution []float64
layer_defs []layer_def
Hidden_layer_sizes []int
tdtrainer_options trainer_def

}


func new_brain_opt () brain_opt {

var this brain_opt

this.temporal_window = undefined_i
this.experience_size = undefined_i
this.start_learn_threshold = undefined_f
this.gamma = undefined_f
this.learning_steps_total = undefined_i
this.learning_steps_burnin = undefined_i
this.epsilon_min = undefined_f
this.epsilon_test_time = undefined_f
this.random_action_distribution = nil
//this.layer_defs = []layer_def{}
this.Hidden_layer_sizes =  []int{}
this.tdtrainer_options = new_trainer_def()

return this

}

type Brain struct {
temporal_window int
experience_size int
start_learn_threshold float64
gamma float64
learning_steps_total int
learning_steps_burnin int
epsilon_min float64
epsilon_test_time float64
random_action_distribution []float64
net_inputs int
num_states int
num_actions int
window_size int
state_window [][]float64
action_window []float64
reward_window []float64
net_window [][]float64
value_net *Net
forward_passes int
last_input_array []float64
learning bool
epsilon float64
age float64
latest_reward float64
average_reward_window *win
average_loss_window *win
experience []*deepqlearn
tdtrainer Trainer

}

func (this *Brain) setup(num_states int, num_actions int, opt brain_opt) {

//var opt = opt || {};
// in number of time steps, of temporal memory
// the ACTUAL input to the net will be (x,a) temporal_window times, and followed by current x
// so to have no information from previous time step going into value function, set to 0.
//this.temporal_window = typeof opt.temporal_window !== 'undefined' ? opt.temporal_window : 1;

if (opt.temporal_window) != undefined_i {

this.temporal_window = opt.temporal_window
} else {
this.temporal_window = 1
}



// size of experience replay memory
//this.experience_size = typeof opt.experience_size !== 'undefined' ? opt.experience_size : 30000;

if (opt.experience_size) != undefined_i {
this.experience_size = opt.experience_size
} else {
this.experience_size = 30000
}

// number of examples in experience replay memory before we begin learning
//this.start_learn_threshold = typeof opt.start_learn_threshold !== 'undefined'? opt.start_learn_threshold : Math.floor(Math.min(this.experience_size*0.1, 1000));
if (opt.start_learn_threshold) != undefined_f {
this.start_learn_threshold = opt.start_learn_threshold
} else {
this.start_learn_threshold = 30000
}

// gamma is a crucial parameter that controls how much plan-ahead the agent does. In [0,1]
//this.gamma = typeof opt.gamma !== 'undefined' ? opt.gamma : 0.8;
if (opt.gamma) != undefined_f {
this.gamma = opt.gamma
} else {
this.gamma = 0.8
}
// number of steps we will learn for
//this.learning_steps_total = typeof opt.learning_steps_total !== 'undefined' ? opt.learning_steps_total : 100000;
if (opt.learning_steps_total) != undefined_i {
this.learning_steps_total = opt.learning_steps_total
} else {
this.learning_steps_total = 100000
}
// how many steps of the above to perform only random actions (in the beginning)?
//this.learning_steps_burnin = typeof opt.learning_steps_burnin !== 'undefined' ? opt.learning_steps_burnin : 3000;
if (opt.learning_steps_burnin) != undefined_i {
this.learning_steps_burnin = opt.learning_steps_burnin
} else {
this.learning_steps_burnin = 3000
}
// what epsilon value do we bottom out on? 0.0 => purely deterministic policy at end
//this.epsilon_min = typeof opt.epsilon_min !== 'undefined' ? opt.epsilon_min : 0.05;

if (opt.epsilon_min) != undefined_f {
this.epsilon_min = opt.epsilon_min
} else {
this.epsilon_min = 0.05
}


// what epsilon to use at test time? (i.e. when learning is disabled)
//this.epsilon_test_time = typeof opt.epsilon_test_time !== 'undefined' ? opt.epsilon_test_time : 0.01;

if (opt.epsilon_test_time) != undefined_f {
this.epsilon_test_time = opt.epsilon_test_time
} else {
this.epsilon_test_time = 0.01
}


// advanced feature. Sometimes a random action should be biased towards some values
// for example in flappy bird, we may want to choose to not flap more often
if ( opt.random_action_distribution != nil) {
// this better sum to 1 by the way, and be of length this.num_actions
//reflect.DeepEqual(this.random_action_distribution,opt.random_action_distribution);
if (len(this.random_action_distribution) != num_actions) {
println("TROUBLE. random_action_distribution should be same length as num_actions.");
}
var a = this.random_action_distribution;
var s = (0.0); for k := 0; k < len(a); k++ {
s +=  (a[k]); }
if (math.Abs(s - 1.0) > 0.0001) {
println("TROUBLE. random_action_distribution should sum to 1!"); }
} else {
this.random_action_distribution = make([]float64, 0);
}

// states that go into neural net to predict optimal action look as
// x0,a0,x1,a1,x2,a2,...xt
// this variable controls the size of that temporal window. Actions are
// encoded as 1-of-k hot vectors
this.net_inputs = num_states * this.temporal_window + num_actions * this.temporal_window + num_states;
this.num_states = num_states;
this.num_actions = num_actions;


this.window_size = int(math.Max(float64(this.temporal_window), 2.0)); // must be at least 2, but if we want more context even more
//this.state_window []float64;
this.state_window =  make([][]float64, this.window_size)
//for x := 0; x < (this.window_size); x++ {
//	yf := make([]float64, 0)
//	for y := 0; y < len(this.state_window); y++ {
//		yf = append(yf, (0.0))
//	}
//
//	this.state_window = append(this.state_window, yf)
//
//}
////this.action_window = []float64{};
this.action_window =  make([]float64,this.window_size)
//for x := 0; x < (1); x++ {
//	this.action_window = append(this.action_window, (float64(this.window_size)))
//}
////this.reward_window = []float64{};
this.reward_window =  make([]float64, this.window_size)
//for x := 0; x < (1); x++ {
//	this.reward_window = append(this.reward_window, (float64(this.window_size)))
//}
//
//
////this.net_window = []float64{};
this.net_window =  make([][]float64, this.window_size)
//for x := 0; x < (this.window_size); x++ {
//	yf := make([]float64, 0)
//	for y := 0; y < len(this.net_window); y++ {
//		yf = append(yf, (0.0))
//	}
//
//	this.net_window = append(this.net_window, yf)
//
//}

// create [state -> value of all possible actions] modeling net for the value function
var layer_defs []layer_def
if (opt.layer_defs != nil) {
// this is an advanced usage feature, because size of the input to the network, and number of
// actions must check out. This is not very pretty Object Oriented programming but I can't see
// a way out of it :(

layer_defs = opt.layer_defs;
if (len(layer_defs) < 2) {
println("TROUBLE! must have at least 2 Layers"); }
if (layer_defs[0].Type != "input") {
println("TROUBLE! first layer must be input layer!"); }
if (layer_defs[len(layer_defs) - 1].Type != "regression") {
println("TROUBLE! last layer must be input regression!"); }
if (layer_defs[0].Out_depth * layer_defs[0].Out_sx * layer_defs[0].Out_sy != this.net_inputs) {
println("TROUBLE! Number of inputs must be num_states * temporal_window + num_actions * temporal_window + num_states!");
}



if (layer_defs[len(layer_defs) - 1].num_neurons != this.num_actions) {
println("TROUBLE! Number of regression neurons should be num_actions!");
}
} else {


// create a very simple neural net by default
ld := new_layer_def()
ld.Type = "input"
ld.Out_sx = 1
ld.Out_sy = 1
ld.Out_depth = this.net_inputs

layer_defs = append(layer_defs, ld);
if(opt.Hidden_layer_sizes != nil) {
// allow user to specify this via the option, for convenience
hl  := opt.Hidden_layer_sizes;
for k:=0;k<len(hl);k++ {
ld2 := new_layer_def()
ld2.Type = "fc"
ld2.num_neurons = hl[k]
ld2.activation = "relu"
layer_defs = append(layer_defs, ld2); // relu by default
}
}

ld3 := new_layer_def()
ld3.Type = "regression"
ld3.num_neurons = num_actions
layer_defs = append(layer_defs, ld3); // value function output

}


this.value_net = new(Net);
this.value_net.makeLayers(layer_defs);

// and finally we need a Temporal Difference Learning trainer!
var tdtrainer_options = new_trainer_def()
tdtrainer_options.learning_rate =0.01
tdtrainer_options.momentum =0.0
tdtrainer_options.batch_size = 64
tdtrainer_options.l2_decay =0.01

if( !reflect.DeepEqual(trainer_def{}, opt.tdtrainer_options)) {

tdtrainer_options = opt.tdtrainer_options; // allow user to overwrite this
}



this.tdtrainer = new_trainer_with_defs(this.value_net, tdtrainer_options);

// experience replay
this.experience = make([]*deepqlearn, 0);
//this.experience  = new(deepqlearn)

// various housekeeping variables
this.age = 0; // incremented every backward()
this.forward_passes = 0; // incremented every forward()
this.epsilon = 1.0; // controls exploration exploitation tradeoff. Should be annealed over time
this.latest_reward = 0;
this.last_input_array = []float64{};
this.average_reward_window = new_window(1000, 10);
this.average_loss_window = new_window(1000, 10);
this.learning = true;




}


func (this *Brain) random_action() int {

// a bit of a helper function. It returns a random action
// we are abstracting this away because in future we may want to
// do more sophisticated things. For example some actions could be more
// or less likely at "rest"/default state.
if(len(this.random_action_distribution) == 0) {
return randi(0, this.num_actions);
} else {
// okay, lets do some fancier sampling:
var p = randf(0, 1.0);
var cumprob = 0.0;
for k:=0;k<this.num_actions;k++ {
cumprob += (this.random_action_distribution[k]);
if(p < cumprob) { return k; }
}
}

return undefined_i
}


type policy struct {

action int
value float64

}

func (this *Brain) policy(s []float64) policy {

// compute the value of doing any action in this state
// and return the argmax action and its value
var sVol = New_Vol(1, 1, this.net_inputs, undefined_f);
sVol.w = s;


//println(len(s),this.net_inputs)


var action_values = this.value_net.forward(sVol, false);
var maxk = 0;
var maxval = (action_values.w[0]);
for k:=1;k<this.num_actions;k++ {
if((action_values.w[k]) > float64(maxval)) { maxk = k; maxval = (action_values.w[k]); }
}

//println((action_values.w[0]),(action_values.w[1]), maxval)

return policy{action:maxk, value:maxval};

}

func (this *Brain) getNetInput(xt []float64) []float64 {


// return s = (x,a,x,a,x,a,xt) state vector.
// It's a concatenation of last window_size (x,a) pairs and current state x
var w = make([]float64, 0);
w = append(w, (xt)...); // start with current state
// and now go backwards and append states and actions from history temporal_window times
var n int = this.window_size;
for  k:=0;k<this.temporal_window;k++ {
// state
w = append(w, (this.state_window[n-1-k])...);
// action, encoded as 1-of-k indicator vector. We scale it up a bit because
// we dont want weight regularization to undervalue this information, as it only exists once
var action1ofk []float64
action1ofk = make([]float64 ,this.num_actions);


for q:=0;q<this.num_actions;q++ { action1ofk[q] = 0.0};
//println(int((this.action_window[n-1-k])), len(action1ofk), "test")
action1ofk[int((this.action_window[n-1-k]))] = 1.0*float64(this.num_states);
//for x := 0; x < len(action1ofk); x++ {
//	println((float64(action1ofk[x])), "test")
//	w = append(w,(float64(action1ofk[x])));
//}

w = append(w,action1ofk...);
}


//println(len(w))
return (w);
}


func (this *Brain) forward(input_array []float64) float64 {

// compute forward (behavior) pass given the input neuron signals from body
this.forward_passes += 1;



this.last_input_array = (input_array); // back this up
//copy(this.last_input_array, input_array)

//println(len(this.last_input_array))


// create network input
var action int;
var net_input []float64
if(this.forward_passes > this.temporal_window) {
// we have enough to actually do something reasonable
net_input = this.getNetInput((input_array));

//println(len(net_input) , "get net input")






if(this.learning) {

// compute epsilon for the epsilon-greedy policy
this.epsilon = math.Min(1.0, math.Max(this.epsilon_min, 1.0-(this.age - float64(this.learning_steps_burnin))/(float64(this.learning_steps_total) - float64(this.learning_steps_burnin))));

//println(1.0-(this.age - float64(this.learning_steps_burnin))/(float64(this.learning_steps_total) - float64(this.learning_steps_burnin)) , "test")
} else {
this.epsilon = this.epsilon_test_time; // use test-time value
}
var rf = randf(0,1);


if(rf < this.epsilon) {
// choose a random action with epsilon probability
action = this.random_action();
} else {
// otherwise use our policy to make decision




var maxact = this.policy((net_input));
action = maxact.action;
}
} else {
// pathological case that happens first few iterations
// before we accumulate window_size inputs
net_input = []float64{};
action = this.random_action();
}

// remember the state and action we took for backward pass


this.net_window = this.net_window[1:];



this.net_window = append(this.net_window, (net_input));



this.state_window = this.state_window[1:];
this.state_window = append(this.state_window,(input_array));
_,this.action_window = shift_float(this.action_window);
this.action_window=  append(this.action_window, float64(action));

return float64(action);

}

type ystruct struct {
dim int
val float64
}

func (this *Brain) backward(reward float64) bool {




this.latest_reward = reward;
this.average_reward_window.add(reward);
//println(len(this.reward_window))
_,this.reward_window =  shift_float(this.reward_window);
//println(len(this.reward_window))
this.reward_window = append(this.reward_window, (reward));


//println(this.reward_window[len(this.reward_window)-1])



if(!this.learning) { return false; }

// various book-keeping
this.age += 1;



// it is time t+1 and we have to store (s_t, a_t, r_t, s_{t+1}) as new experience
// (given that an appropriate number of state measurements already exist, of course)
if(this.forward_passes > this.temporal_window + 1) {

	var e = new_experience([]float64{},undefined_f,undefined_f,[]float64{}); // this said new Experience, this is a little worry
	var n = this.window_size;
	e.state0 = (this.net_window[n-2]);
	e.action0 = (this.action_window[n-2]);

	//println((e.action0), "action0")



	e.reward0 = (this.reward_window[n-2]);

	//println((e.reward0), "reward")

	e.state1 = this.net_window[n-1];


	//println(len(this.net_window), "win")


	if(len(this.experience) < this.experience_size) {
		this.experience = append(this.experience, e);
	} else {
		// replace. finite memory!
		var ri = randi(0, this.experience_size);
		this.experience[ri] = e;
	}





}

	// learn based on experience, once we have some samples to go on
	// this is where the magic happens...



	if(len(this.experience) > int(this.start_learn_threshold)) {

		//println("test")
		var avcost = 0.0;
		for k:=0;k < this.tdtrainer.batch_size;k++ {


			var re = randi(0, len(this.experience));




			var e = this.experience[re];

			//println(len(e.state1) ,"test")
			var x = New_Vol(1, 1, this.net_inputs, undefined_f);


			x.w = e.state0;



			var maxact = this.policy(e.state1);
			var r = (e.reward0) + this.gamma * maxact.value;

			//println(e.reward0)
			var ystruct ystruct = ystruct{dim: int((e.action0)), val: r};

			var loss = this.tdtrainer.train(x, []float64{}, ystruct); // maybe  []float64{r} is wrong


			//println(loss)

			avcost += loss; // loss.loss
		}
		avcost = (avcost/float64(this.tdtrainer.batch_size));
		this.average_loss_window.add(avcost);
	}



	return false

}


type deepqlearn struct {

	state0 []float64
	action0 float64
	reward0 float64
	state1 []float64
	Brain Brain

}



func new_experience (state0 []float64, action0 float64, reward0 float64, state1 []float64) *deepqlearn {

	var this *deepqlearn = new(deepqlearn)

	this.state0 = (state0);
	this.action0 = (action0);
	this.reward0 = (reward0);
	this.state1 = (state1);

	return this
}





func shift_float(a []float64) (float64, []float64) {
	x, a := a[0], a[1:]

	return x, a
}

//func shift_mfloat(a []float64) (float64, []float64) {
//	x, a := a[0], a[1:]
//
//	return (x), a
//}

//func shift_int(a []int) (int, []int) {
//	x, a := a[0], a[1:]
//
//	return x, a
//}
//
//func shift_mint(a []int) (int, []int) {
//	x, a := a[0], a[1:]
//
//	return get_int(x), a
//}

//func (f []float64) []float64 {
//
//
//
//	var mf = make([]float64, 0)
//
//	for x := 0; x < len(f); x++ {
//		mf = append(mf, (f[x]))
//	}
//
//	return mf
//
//}

//func mfloat_arr_to_float_arr(f []float64) []float64 {
//
//	var mf = make([]float64, 0)
//
//	for x := 0; x < len(f); x++ {
//		mf = append(mf, (f[x]))
//	}
//
//	return mf
//
//}


//func int_arr_to_mint_arr(f []int) []int {
//
//	var mf = make([]int, 0)
//
//	for x := 0; x < len(f); x++ {
//		mf = append(mf, set_mint(f[x]))
//	}
//
//	return mf
//
//}

//func mint_arr_to_int_arr(f []int) []int {
//
//	var mf = make([]int, 0)
//
//	for x := 0; x < len(f); x++ {
//		mf = append(mf, get_int(f[x]))
//	}
//
//	return mf
//
//}
//
//func mint_to_float(f int) float64 {
//
//	return float64(get_int(f))
//
//}
//
//func mfloat_to_int(f *float64) int {
//
//	return int((f))
//
//}


//func merge_nets (net []*Net) *Net {
//
//
//	net0 := new(Net)
//
//	net0.fromJSON(net[0].toJSON())
//
//
//
//if len(net) == 1{
//	return net[0]
//}
//
//
//
//	total_scores := make([]float64, len(net))
//	total := 0.0
//	for x := 0; x < len(net); x++ {
//
//		for y := 0 ; y < len(net[x].score); y++ {
//	total_scores[x] += (net[x].score[y]) / (net[x].trained_epochs - float64(y))
////println(1 -net[x].score[y] )
//
//		}
//		total_scores[x] /= net[x].trained_epochs
//		total += total_scores[x]
//
//	}
//
//
//
//
//for y:= 0; y < len(net[0].Layers); y++ {
//
//
//	for t := 0; t < len(net[0].Layers[y].l1_decay_mul); t++ {
//
//		if net[0].Layers[y].l1_decay_mul != nil {
//
//			diff := make([]float64, len(net[0].Layers[y].l1_decay_mul))
//
//
//			for x := 0; x < len(net); x++ {
//
//				// add up w's and Dw's
//				for q := 0; q < len(net[0].Layers[y].l1_decay_mul); q ++ {
//
//					diff[q] += (net[x].Layers[y].l1_decay_mul[q]) *  (total_scores[x])
//
//				}
//
//
//
//			}
//
//
//			//	println((net[0].Layers[y].filters[t][j].w[0]) , "<--")
//			//println((net[1].Layers[y].filters[t][j].w[0]) , "<--")
//			//println((net[2].Layers[y].filters[t][j].w[0]) , "<--")
//
//
//			for q := 0; q < len(net[0].Layers[y].l1_decay_mul); q ++ {
//				net0.Layers[y].l1_decay_mul[q] = ( diff[q]  / total)
//				//net[0].Layers[y].filters[t][j].Dw[q] = ((net[0].Layers[y].filters[t][j].Dw[q]) / float64(len(net)))
//
//			}
//
//		}
//	}
//
////====
//	for t := 0; t < len(net[0].Layers[y].l2_decay_mul); t++ {
//
//		if net[0].Layers[y].l2_decay_mul != nil {
//
//			diff := make([]float64, len(net[0].Layers[y].l2_decay_mul))
//
//
//			for x := 0; x < len(net); x++ {
//
//				// add up w's and Dw's
//				for q := 0; q < len(net[0].Layers[y].l2_decay_mul); q ++ {
//
//					diff[q] += (net[x].Layers[y].l2_decay_mul[q]) *  (total_scores[x])
//
//				}
//
//
//
//			}
//
//
//			//	println((net[0].Layers[y].filters[t][j].w[0]) , "<--")
//			//println((net[1].Layers[y].filters[t][j].w[0]) , "<--")
//			//println((net[2].Layers[y].filters[t][j].w[0]) , "<--")
//
//
//			for q := 0; q < len(net[0].Layers[y].l2_decay_mul); q ++ {
//				net0.Layers[y].l2_decay_mul[q] = ( diff[q]  / total)
//				//net[0].Layers[y].filters[t][j].Dw[q] = ((net[0].Layers[y].filters[t][j].Dw[q]) / float64(len(net)))
//
//			}
//
//		}
//	}
//
//
//	for t := 0; t < len(net[0].Layers[y].biases); t++ {
//
//		if net[0].Layers[y].biases[t] !=nil {
//
//		diff := make([]float64, len(net[0].Layers[y].biases[t].w))
//
//
//		for x := 0; x < len(net); x++ {
//
//			// add up w's and Dw's
//			for q := 0; q < len(net[0].Layers[y].biases[t].w); q ++ {
//
//				diff[q] += (net[x].Layers[y].biases[t].w[q]) *  (total_scores[x])
//
//			}
//
//
//			//for q := 0; q < len(net[0].Layers[y].l1_decay_mul); q ++ {
//			//
//			//}
//		}
//
//
//		//	println((net[0].Layers[y].filters[t][j].w[0]) , "<--")
//		//println((net[1].Layers[y].filters[t][j].w[0]) , "<--")
//		//println((net[2].Layers[y].filters[t][j].w[0]) , "<--")
//
//
//		for q := 0; q < len(net[0].Layers[y].biases[t].w); q ++ {
//			net0.Layers[y].biases[t].w[q] = ( diff[q]  / total)
//			//net[0].Layers[y].filters[t][j].Dw[q] = ((net[0].Layers[y].filters[t][j].Dw[q]) / float64(len(net)))
//
//		}
//
//	}
//	}
//for t := 0; t < len(net[0].Layers[y].filters); t++ {
//
//				for j := 0; j < len(net[0].Layers[y].filters[t]); j++ {
//
//
//
//
//
//					diff := make([]float64, len(net[0].Layers[y].filters[t][j].w))
//
//
//					for x := 0; x < len(net); x++ {
//
//						// add up w's and Dw's
//						for q := 0; q < len(net[0].Layers[y].filters[t][j].w); q ++ {
//
//							diff[q] += (net[x].Layers[y].filters[t][j].w[q]) *  (total_scores[x])
//
//						}
//
//
//						for q := 0; q < len(net[0].Layers[y].l1_decay_mul); q ++ {
//
//						}
//					}
//
//
//				//	println((net[0].Layers[y].filters[t][j].w[0]) , "<--")
//					//println((net[1].Layers[y].filters[t][j].w[0]) , "<--")
//					//println((net[2].Layers[y].filters[t][j].w[0]) , "<--")
//
//
//					for q := 0; q < len(net[0].Layers[y].filters[t][j].w); q ++ {
//						net0.Layers[y].filters[t][j].w[q] = ( diff[q]  / total)
//
//						//net[0].Layers[y].filters[t][j].Dw[q] = ((net[0].Layers[y].filters[t][j].Dw[q]) / float64(len(net)))
//
//					}
//
//					//for q := 0; q < len(net0.Layers[y].l1_decay_mul); q ++ {
//					//	//net[0].Layers[y].l1_decay_mul[q] /= float64(len(net) )
//					//	//net[0].Layers[y].l2_decay_mul[q] /= float64(len(net) )
//					//	//net[0].Layers[y].filters[t][j].Dw[q]  =  ((net[0].Layers[y].filters[t][j].Dw[q])  /  float64(len(net) ))
//					//}
//
//
//
//				//	println((net[0].Layers[y].filters[t][j].w[0]))
//
//
//				}
//			}
//
//
//
//	}
//
//
//
//
//return net0
//}


func make_net_group (layer_def []layer_def, N int) []*Net {


	var net []*Net = make([]*Net, 0);



	for x := 0; x < N; x++ {

		net = append(net, new(Net))
		net[x].makeLayers(layer_def);
	}

	return net


	if N == 1 {
		return net
	}

	for x := 1; x < len(net); x++ {





		for y:= 0; y < len(net[0].Layers); y++ {




			for t := 0; t < len(net[0].Layers[y].filters); t++ {

				for j := 0; j < len(net[0].Layers[y].filters[t]); j++ {



					// add up w's and Dw's
					for q := 0; q < len(net[0].Layers[y].filters[t][j].w); q ++ {
						net[x].Layers[y].filters[t][j].w[q]  = ((net[0].Layers[y].filters[t][j].w[q]))
						//net[x].Layers[y].filters[t][j].Dw[q]  = ((net[0].Layers[y].filters[t][j].Dw[q]))
					}


					//for q := 0; q < len(net[0].Layers[y].l1_decay_mul); q ++ {
					//	net[0].Layers[y].l1_decay_mul[q] += net[x].Layers[y].l1_decay_mul[q]
					//	net[0].Layers[y].l2_decay_mul[q] += net[x].Layers[y].l2_decay_mul[q]
					//}






				}
			}




			for t := 0; t < len(net[0].Layers[y].biases); t++ {

				if net[0].Layers[y].biases[t] !=nil {




					for x := 0; x < len(net); x++ {

						// add up w's and Dw's
						for q := 0; q < len(net[0].Layers[y].biases[t].w); q ++ {

							net[x].Layers[y].biases[t].w[q] =  ((net[0].Layers[y].biases[t].w[q]))


						}



					}




				}
			}


			for t := 0; t < len(net[0].Layers[y].l2_decay_mul); t++ {

				if net[0].Layers[y].l2_decay_mul != nil {




					for x := 0; x < len(net); x++ {

						// add up w's and Dw's
						for q := 0; q < len(net[0].Layers[y].l2_decay_mul); q ++ {

							net[x].Layers[y].l2_decay_mul[q]  = net[0].Layers[y].l2_decay_mul[q]

						}



					}




				}
			}



			if net[0].Layers[y].l1_decay_mul != nil {




				for x := 0; x < len(net); x++ {

					// add up w's and Dw's
					for q := 0; q < len(net[0].Layers[y].l1_decay_mul); q ++ {

						net[x].Layers[y].l1_decay_mul[q]  = net[0].Layers[y].l1_decay_mul[q]

					}



				}




			}




		}

	}


	return net

}


// extract system
type Changeable interface {
	Set(x, y int, c color.Color)
	ColorModel() color.Model
	At(x, y int) color.Color
}


func memset(arr []int, offset int, length int, value int) []int {
	for  i := 0; i < length; i++ {

		arr[offset] = value;
		offset++

	}

	return arr
};

type blb_a struct {
	T int
	q int
}

//func assert(exp, message) {
//if (!exp) {
//throw new AssertException(message);
//}
//}

type blb_dta struct {

	data []int
	pos []int
	label []int
	c int
	traced bool
	S int
	offset int
	max int

}

func tracer(data []int, pos []int, label []int, max int, S int, p int,BACKGROUND int, FOREGROUND int, UNSET int, MARKED int) (bool, blb_a, blb_dta)  {

	//println("test",S,p)
	for  d := 0; d < 8; d++ {

		var T int
		var q int
		q = (p + d) % 8;

		if q >= 0 && q < len(pos) {
			T = S + pos[q];
		} else {
			T = S
		}
		//println(S)
		// Make sure we are inside image
		if T < 0 || T >= max {
			//continue;
		} else {

			if (T >= 0 && T < len(data) && data[T] != BACKGROUND  ) {

				return true, blb_a{T:T, q:q}, blb_dta{label: label, data:data, pos: pos};
			}

			//assert(label[T] <= UNSET); // this line seems to not be necissary


			label[T] = MARKED;

		}
	}


	// No move
	return false, blb_a{T:S, q:-1}, blb_dta{label:label, data:data, pos: pos};
};

/**
	 *
	 * @param S Offset of starting point
	 * @param C label count
	 * @param external Boolean Is this internal or external tracing
	 */
func contourTracing(data []int, pos []int, label []int, max int, S int, C int, external bool,BACKGROUND int, FOREGROUND int, UNSET int, MARKED int) blb_dta {

	//println(S ,"test")
	var p int
	if external {
		p = 7
	} else {
		p = 3
	}


	// Find out our default next pos (from S)
	_, tmp, blbdta := tracer(data, pos, label, max, S, p, BACKGROUND , FOREGROUND, UNSET, MARKED);
	label = blbdta.label
	//data = blbdta.data
	//pos  = blbdta.pos
	var T2 int
	var q int

	//println(S, "s")


	T2  = tmp.T;
	q = tmp.q;


	label[S] = C;

	//println(T2, S)

	// Single pixel check
	if (T2 == S) {
		//println("huh")
		return blb_dta{data:data,pos:pos,label:label, c:C };
	}

	var counter = 0;

	var Tnext  int = T2;
	var T    int   = T2;



	for ok := true; ok; ok =  ok  {

		counter++
		//if counter < max {
		//ok = false
		//assert(counter++ < max, "Looped too many times!"); // not necissary to say something
		//return blb_dta{data:data,pos:pos,label:label, c:C }
		//println("Looped too many times!")
		//} else {



		if Tnext < len(label) && Tnext >= 0 {
			label[Tnext] = C;
			//println(label[Tnext], "nat")
		}



		T = Tnext;
		p = (q + 5) % 8 ;

		//println(Tnext, "Tnext")
		var blbdta2 blb_dta
		//println("test44s")
		_, tmp, blbdta2 := tracer(data, pos, label, max, T, p, BACKGROUND , FOREGROUND, UNSET, MARKED);
		label = blbdta2.label



		Tnext = tmp.T;
		q     = tmp.q;

		//fmt.Println(Tnext != T2, T != S)
		if ( T != S ) == false || (Tnext != T2 ) == false  {
			ok = false
		}

	}
	//}

	return blb_dta{data:data,pos:pos,label:label, c:C }
};


func extract(data []int, pos []int, label []int, c int, max int, h int, w int, BACKGROUND int, FOREGROUND int, UNSET int, MARKED int) blb_dta {
	var traced bool
	var y int = 1; // We start at 1 to avoid looking above the image
	//y++; // check this is ok
	for ok := true; ok; ok =  ok {


		var x = 0;
		//x++; // check this is ok
		for ok2 := true; ok2; ok2 = ok2  {



			var offset = y * w + x;

			//println(offset,y,w,x) // got that part
			// this passes

			// We skip white pixels or previous labeled pixels
			if offset < len(data) && data[offset] == BACKGROUND  {

				//continue;
			} else {


				traced = false;



				// Step 1 - P not labelled, and above pixel is white
				if (offset - w < len(data) && data[offset - w] == BACKGROUND && offset < len(label) && label[offset] == UNSET ) {

					//console.log(x + "," + y + " step 1");

					// P must be external contour
					//println(offset, c, w)
					var blbdta blb_dta = contourTracing(data , pos , label ,max,  offset, c, true, BACKGROUND , FOREGROUND, UNSET, MARKED);
					//println(c, 2)

					data = blbdta.data;
					pos = blbdta.pos;
					label = blbdta.label;
					//c = blbdta.c
					//println(c)
					c++;

					traced = true;
				}

				// Step 2 - Below pixel is white, and unmarked
				if (offset + w < len(data) &&  offset + w < len(label) &&  data[offset + w] == BACKGROUND && label[offset + w] == UNSET) {
					//console.log(x + "," + y + " step 2");
					//println(c, "huhuh");
					// Use previous pixel label, unless this is already labelled
					var n = label[offset - 1];
					if (label[offset] != UNSET) {
						n = label[offset];
					}
					//if (n > UNSET) {
					//ok2 = false
					//ok = false
					//	println("assert1");
					//} else {
					//assert( n > UNSET, "Step 2: N must be set, (" + x + "," + y + ") " + n + " " + data[offset - 1]);

					// P must be a internal contour
					var blbdta blb_dta  = contourTracing(data , pos , label ,max, offset, n, false, BACKGROUND , FOREGROUND, UNSET, MARKED);




					data = blbdta.data;
					pos = blbdta.pos;
					label = blbdta.label;
					//c = blbdta.c
					//println("wtf")
					traced = true;


					//}
				}

				// Step 3 - Not dealt with in previous two steps
				if offset <  len(label) && label[offset] == UNSET {
					//console.log(x + "," + y + " step 3");
					//console.log2D(label, w, h);
					var n int = label[offset - 1];

					//if (!traced) {
					//ok2 = false
					//ok = false
					//println("assert2");
					//} else {

					//if (n > UNSET) {
					//ok2 = false
					//ok = false
					//	println("assert3");
					//} else {

					// Assign P the value of N

					label[offset] = n;

					//}
					//}
				}

			}


			ok2 = (x+1 < w)
			x++

		} ;


		ok = (y+1 < (h-1))
		y++

	} ; // We end one before the end to to avoid looking below the image

	//println("labels=", c);
	return blb_dta{data:data,pos:pos,label:label, c:c, traced:traced };
}

func BlobExtraction(data []int, w int, h int, BACKGROUND int, FOREGROUND int, UNSET int,  MARKED int) []int {
	var max int = w * h;



	/*
	 * 5 6 7
	 * 4 P 0
	 * 3 2 1
	 */
	var pos []int = []int{1, w + 1, w, w -1, -1, -w -1, -w, -w+1}; // Clockwise

	var label = make([]int,max); // Same size as data
	var c = 1;      // Component index

	// We change the border to be white. We could add a pixel around
	// but we are lazy and want to do this in place.
	// Set the outer rows/cols to min
	data = memset(data, 0, w, BACKGROUND); // Top
	data =  memset(data, w * (h-1), w, BACKGROUND); // Bottom

	for y := 1; y < h-1; y++ {
		var offset int = y * w;
		data[offset        ] = BACKGROUND; // Left
		data[offset + w - 1] = BACKGROUND; // Right
	}

	// Set labels to zeros
	label  = memset(label, 0, max, UNSET);




	var blbdta blb_dta = extract(data, pos, label, c, max, h, w, BACKGROUND , FOREGROUND, UNSET, MARKED);



	return blbdta.label
}

type blob_bounds struct {
	l int
	x1 int
	y1 int
	x2 int
	y2 int
	area int
}

func BlobBounds(label []int, width int, height int) []blob_bounds {
	var blob = []blob_bounds{};

	blob = append(blob,blob_bounds{l:0, x1:0, y1:0, x2:0, y2:0, area:0});
	var offset int = 0;
	for  y := 0; y < height; y++ {
		for  x := 0; x < width; x++ {


			//if offset < len(label) {
			var l int = label[offset]
			offset++
			if (l <= 0) {
				//continue;
			} else {

				singleton_test := false
				for i := 1; i < len(blob); i++ {
					if (l == blob[i].l) {
						singleton_test = true
					}
				}

				if singleton_test {


					//var b blob_bounds = blob[l];


					if ( l < len(blob) && blob[l].x2 < x) {

						blob[l].x2 = x;
					}

					if (l < len(blob) &&  blob[l].x1 > x) {

						blob[l].x1 = x;
					}

					// As we are going from top down, the bottom y should increase
					if (l < len(blob) ) {
						blob[l].y2 = y;
					}
					//println(blob[l].x2)

					//blob[l] = blob_bounds{l:l, x1:b.x1, y1:b.y1, x2:b.x2, y2:b.y2}
					//				blob[l] = b;
				} else {
					if (l >= len(blob)) {
						blob = append(blob,blob_bounds{l:l, x1:x, y1:y, x2:x, y2:y})
					} else{

						blob[l] = blob_bounds{l:l, x1:x, y1:y, x2:x, y2:y}
					}
				}

			}

		}
	}
	//}




	for  i := 1; i < len(blob); i++ {


		blob[i].area = (blob[i].x2 - blob[i].x1 + 1) * (blob[i].y2 - blob[i].y1 + 1);
	}

	return blob;
}

/*
UnionFind contains information about how given numbers are related. Use
UnionFind.Union(2, 3) to specify that 2 and 3 belong to the same set.
Additional calls to Union can create relationships by the transitive property.
Query the relationship between numbers with the Connected or Find methods.
This implementation of UnionFind uses quick-find and stores connection data in
an array (implemented as a slice of ints).
*/
type UnionFind struct {
	id    []int
	count int
}

// New initializes a new UnionFind struct and returns a pointer.
func NewUnionFind(count int) *UnionFind {
	id := make([]int, count)
	for i := 0; i < count; i++ {
		id[i] = i
	}
	return &UnionFind{id: id, count: count}
}

// Find returns the set ID of a given element.
func (uf UnionFind) Find(p int) int {
	return uf.id[p]
}

// Connected determines whether two elements belong to the same set.
func (uf UnionFind) Connected(p, q int) bool {
	return uf.id[p] == uf.id[q]
}

// Union specifies two elements as belonging to the same set for future query
// with the Connected or Find methods.
func (uf *UnionFind) Union(p, q int) {
	if uf.Connected(p, q) {
		return
	}
	p_id := uf.id[p]
	for i := 0; i < len(uf.id); i++ {
		if uf.id[i] == p_id {
			uf.id[i] = uf.id[q]
		}
	}
	uf.count--
}


func  imgSave(img Changeable, filename string) {
	//for _, row := range c.contents {

	toimg, _ := os.Create(filename)
	defer toimg.Close()

	_ = png.Encode(toimg, img.(image.Image))
	//}
}

func get_bg_color(img_arr []int, w int, h int) int {

	bg := 255;

	//println(len(img_arr))

	bg_col := []int{img_arr[0], img_arr[w-1], img_arr[(h-1)*(w-1)-(w-1)], img_arr[(h-1)*(w-1)]}
	bg_col_count := []int{0,0,0,0}

	for z := 0; z < len(bg_col); z++ {
		for z2 := 0; z2 < len(bg_col); z2++ {
			if  z2 != z && math.Abs(float64(bg_col[z] - bg_col[z2])) <= 25.0  {
				bg_col_count[z]++
			}
		}

	}
	best_label := 255
	best_score := 0

	for z := 0; z < len(bg_col_count); z++ {

		if best_score < bg_col_count[z] {
			best_score = bg_col_count[z]
			best_label = bg_col[z]
		}

	}


	if best_score < 1 {
		// cycle through the labels

		for y := 0; y < h; y++ {
			for x := 0; x < w; x++ {

				for z := 0; z < len(bg_col_count); z++ {

					if math.Abs(float64(bg_col[z] - img_arr[y*w+x])) <= 25.0 {
						bg_col_count[z]++

					}


				}

			}
		}

		best_label = 255
		best_score = 0

		for z := 0; z < len(bg_col_count); z++ {

			if best_score < bg_col_count[z] {
				best_score = bg_col_count[z]
				best_label = bg_col[z]
			}

		}

		bg = best_label


	} else {

		bg = best_label

	}
	return bg

}

func remove_smaller_objects(arr []int, w,h, background_variation int ) []int{

	arr2 := []int{};

	for x := 0; x < len(arr); x++ {

		arr2 = append(arr2, arr[x]);

	}
	//
	//for y := 0; y < size.Y; y++	{
	//	for x := 0; x < size.X; x++ {
	//
	//		r,g,b,_ := img.At(x,y).RGBA()
	//
	//
	//
	//		ans := int(float64(r+g+b) / 3 / 65535 * 255)
	//
	//		// get the real bg color
	//		arr2 = append(arr2, ans)
	//
	//
	//
	//		//print(ans,",");
	//		arr = append(arr, ans)
	//
	//	}
	//	//println("")
	//}





	real_bg := get_bg_color(arr2, w, h)


	for x := 0; x < len(arr); x++ {

		if math.Abs(float64(arr[x] - real_bg)) <= float64(background_variation)  {
			//ans = 0

			arr[x] = 255
		} else
		{
			arr[x] = 0
		}

	}

	//These are constants
	var BACKGROUND int = 255;
	var FOREGROUND int = 0;
	var UNSET int      = 0;
	var MARKED int    = -1;

	BACKGROUND = get_bg_color(arr, w, h)




	//fmt.Println("bg=", real_bg)

	bb := BlobExtraction(arr,w, h, BACKGROUND , FOREGROUND, UNSET, MARKED);

	labels := map[int]int{};
	//label_touching_left := map[int]bool{};
	//label_touching_right := map[int]bool{};

	//largest_num := 0
	//lagest_label := 0


	//larger_than := map[int]bool{}


	for x := 0; x < len(bb); x++ {

		if bb[x] > 0 {
			labels[bb[x]]++



			//if largest_num < labels[bb[x]] {
			//	largest_num = labels[bb[x]]
			//	lagest_label = bb[x]
			//
			//}





			//if labels[bb[x]] > i {
			//larger_than[bb[x]] = true
			//}

		}


	}

	b := int(real_bg)

	for y := 0; y < h; y++	{
		for x := 0; x < w; x++ {

			if bb[y*w+x] == 0 {
				arr[y*w+x] = b
				//chng.Set(x, y, color.RGBA{R:b, G:b, B:b, A:255})
			} else if bb[y*w+x] == -1 {
				arr[y*w+x] = b

				//chng.Set(x,y,color.RGBA{R:b,G:b,B:b,A:255})
			} else if labels[bb[y*w+x]] < (int(float64(w*h)/4)) {
				arr[y*w+x] = b
				//chng.Set(x,y,color.RGBA{R:b,G:b,B:b,A:255}) // remove all but the largest object outside of bg
			}
			//=else {
			////chng.Set(x, y, color.RGBA{R:b, G:b, B:b, A:255})
			//
			//for z := range bigger_than {
			//
			//if z == bb[y*size.X+x] {
			//
			//chng.Set(x,y,color.RGBA{R:b,G:b,B:b,A:255})
			//
			//}
			//
			//}
			//
			//}




			//print(bb[y*size.X+x],",");

		}
		// println("")
	}


	background_is_same := true
	for x := 0; x < len(arr); x++ {
		if math.Abs(float64(arr[x] - BACKGROUND)) <= float64(background_variation) {
			background_is_same = false
		}
	}

	if background_is_same {
		return arr2;
	} else {
		return arr;
	}


}




func create_net_from_struct(Layers []map[string]interface{}, net *Net ) *Net {




	all_Layers := []layer_def{};

	for x := 0; x < len(Layers); x++ {
		layer := new_layer_def()
		if Layers[x]["type"] != nil {
			layer.Type = Layers[x]["type"].(string)
		}
		if Layers[x]["act"] != nil {
			layer.activation = Layers[x]["act"].(string)
		}

		if (Layers[x]["type"] != "conv" && Layers[x]["type"] != "pool") {
			if Layers[x]["x"] != nil {
				layer.Out_sx = int(Layers[x]["x"].(float64))
			}
			if Layers[x]["y"] != nil {
				layer.Out_sy = int(Layers[x]["y"].(float64))
			}
			if Layers[x]["z"] != nil {
				layer.Out_depth = int(Layers[x]["z"].(float64))
			}
		}  else {
			if Layers[x]["x"] != nil {
				layer.sx = int(Layers[x]["x"].(float64))
			}
			if Layers[x]["y"] != nil {
				layer.sy = int(Layers[x]["y"].(float64))
			}
			if Layers[x]["z"] != nil {
				layer.depth = int(Layers[x]["z"].(float64))
			}
		}
		if Layers[x]["type"] == "regression" {
			if Layers[x]["size"] != nil {
				layer.num_neurons = int(Layers[x]["size"].(float64))
			}
		} else if Layers[x]["type"] == "softmax" {
			if Layers[x]["size"] != nil {
				layer.num_classes = int(Layers[x]["size"].(float64))
			}
		} else if Layers[x]["type"] == "svm" {
			if Layers[x]["size"] != nil {
				layer.num_classes = int(Layers[x]["size"].(float64))
			}
		} else {
			if Layers[x]["size"] != nil {
				layer.num_neurons  = int(Layers[x]["size"].(float64))
			}
		}

		if Layers[x]["filters"] != nil {
			layer.filters	= int(Layers[x]["filters"].(float64))
		}

		if Layers[x]["stride"] != nil {
			layer.stride	= int(Layers[x]["stride"].(float64))
		}

		if Layers[x]["pad"] != nil {
			layer.pad	= int(Layers[x]["pad"].(float64))
		}





		all_Layers = append(all_Layers, layer);

	}

	net.makeLayers(all_Layers);

	return net;
}


func NewLayerDefs () []Layer_def {
	return make([]Layer_def, 0)
}


func New_layer_def() Layer_def {
	return  Layer_def(new_layer_def());
}

func ChangLayerDef(ld Layer_def) Layer_def {

	return ld
}


  // Mat utils
  // fill matrix with random gaussian numbers
  //var fillRandn = func(m int, mu int, std float64) { for(var i=0,n=m.w.length;i<n;i++) { m.w[i] = randn(mu, std); } }
   func fillRand(m *Vol, lo float64, hi float64) {
	    n:=len(m.w)
		for i:=0;i<n;i++ { 
	  m.w[i] = global.randn(lo, hi); 
	} 
}

// return Mat but filled with random numbers from gaussian
 func (this *LSTMNet) RandMat(n int,d int,mu int,std float64) *Vol {
    var m = new(Vol)
	m.Mat(n, d);
    //fillRandn(m,mu,std);
    fillRand(m,-std,std); // kind of :P
    return m;
  }

func InitLSTM(input_size int, Hidden_sizes []int, output_size int) *LSTMNet {
    // Hidden size should be a list

    var this = new(LSTMNet);

this.Model = make(map[string]*Vol)
this.Hidden = make([]*Vol,0)
this.Cell = make([]*Vol,0)
this.Backprop = make([]func(),0)
this.needs_backprop = true

		var Hidden_size int	

    for d:=0;d<len(Hidden_sizes);d++ { // loop over depths
      var prev_size = d
	  if d == 0 {
		prev_size = input_size
	  } else {
		 prev_size = Hidden_sizes[d - 1]
	  }
      Hidden_size = Hidden_sizes[d];

      // gates parameters
      this.Model["Wix"+iToS(d)] = this.RandMat(Hidden_size, prev_size , 0, 0.08);  
      this.Model["Wih"+iToS(d)] = this.RandMat(Hidden_size, Hidden_size , 0, 0.08);
	  mt0 := new(Vol)
	   mt0.Mat(Hidden_size, 1)

      this.Model["bi"+iToS(d)] = mt0;
      this.Model["Wfx"+iToS(d)] = this.RandMat(Hidden_size, prev_size , 0, 0.08);  
      this.Model["Wfh"+iToS(d)] = this.RandMat(Hidden_size, Hidden_size , 0, 0.08);

	   mt := new(Vol)
	   mt.Mat(Hidden_size, 1)

      this.Model["bf"+iToS(d)] = mt;

      this.Model["Wox"+iToS(d)] = this.RandMat(Hidden_size, prev_size , 0, 0.08);  
      this.Model["Woh"+iToS(d)] = this.RandMat(Hidden_size, Hidden_size , 0, 0.08);

	   mt1 := new(Vol)
	   mt1.Mat(Hidden_size, 1)
      this.Model["bo"+iToS(d)] = mt1;
      // Cell write params
     this.Model["Wcx"+iToS(d)] = this.RandMat(Hidden_size, prev_size , 0, 0.08);  
      this.Model["Wch"+iToS(d)] = this.RandMat(Hidden_size, Hidden_size , 0, 0.08);
	    mt2 := new(Vol)
	   mt2.Mat(Hidden_size, 1)
      this.Model["bc"+iToS(d)] = mt2;
    }
    // decoder params
    this.Model["Whd"] = this.RandMat(output_size, Hidden_size, 0, 0.08);
	  mt3 := new(Vol)
	   mt3.Mat(output_size, 1)
    this.Model["bd"] = mt3
    
	return this;
  }



  func Round(val float64, roundOn float64, places int ) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return newVal
}