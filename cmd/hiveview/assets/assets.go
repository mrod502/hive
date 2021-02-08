// Code generated by "esc -pkg assets -o assets/assets.go -prefix assets/ assets/"; DO NOT EDIT.

package assets

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	if !f.isDir {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is not directory", f.name)
	}

	fis, ok := _escDirs[f.local]
	if !ok {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is directory, but we have no info about content of this dir, local=%s", f.name, f.local)
	}
	limit := count
	if count <= 0 || limit > len(fis) {
		limit = len(fis)
	}

	if len(fis) == 0 && count > 0 {
		return nil, io.EOF
	}

	return fis[0:limit], nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// FS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func FS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// Dir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func Dir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// FSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func FSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// FSMustByte is the same as FSByte, but panics if name is not present.
func FSMustByte(useLocal bool, name string) []byte {
	b, err := FSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// FSString is the string version of FSByte.
func FSString(useLocal bool, name string) (string, error) {
	b, err := FSByte(useLocal, name)
	return string(b), err
}

// FSMustString is the string version of FSMustByte.
func FSMustString(useLocal bool, name string) string {
	return string(FSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/app-viewer.js": {
		name:    "app-viewer.js",
		local:   "assets/app-viewer.js",
		size:    4520,
		modtime: 1603658490,
		compressed: `
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
BiRfLf4bAAD//6qI6yuoEQAA
`,
	},

	"/app.js": {
		name:    "app.js",
		local:   "assets/app.js",
		size:    17091,
		modtime: 1612794710,
		compressed: `
H4sIAAAAAAAC/7Q7f3fbuJH/+1PMstmIjCVScpw4li2723hzTZvd7Euy13e1/FKIhCTYFMACoGR34/vs
9wYAf4pynHYvf8QSMBjMDAbzE4oFVxokVXmq1QchNEzAi9z3yNvbyzVLFUzgtz0AgOiZ+QPP4M+ffno3
oDwWCeMLNxiZv0u9Sj+bGTqGec5jzQT3lZaBQ2IQRe+oBr2kcPH+J0gEMA1zISFXYQmzJhISmEAi4nxF
uQ5jSYmmP6YUv/k9TW81kZT0gpNyTRIyzqn8RG+RE6XlSW3L/6GqD297KyAbIimIeW0Ollpn4yhSmsQ3
Yk3lPBWbMBar6J85VciCikYHo+NXL4cRcljyPmB8cE3WRMWSZXpw/c+cyrs64rdwnSsNieA9DWQhKa1Y
lFTnkhdUo0wtvff9vT23nPKYZCpPiaaQEE2AccUSCgQ0sYLXZFETtL7L+vD7S1vfZTvF7Hse7Js9T7Y5
E7nu5KyhSfCD1pLNck2hU6eI1vKrOnW7m3rv1qvRdhsqqssdfW/m9VvUozBImsIEbtsM1JgjaRqqfKa0
ZHzhv+ybgZTyhV7CAF4Guzi+YFLfQS7TQUakYnwBYm5OJ5cpLIlagqILpLshgwXVn3Hyc0YkWamaHOpC
QMol1WuCxP923xg3qgkT2DCeiE2Yipjg+hCx1hgZtQSxJhItgFkdqixl2u89rV86VCUfIRlMYHgCDE7N
IieLE2D7+3UaC8QZYRImBvSSXRWoJ3XUTtprkl4mFI//1w9vX4tVJjgeKyK4HF4FV3jyO6ZHV0GJ7b59
ghb1rnN6bXRIAeFAeLwUckCtPsFcihX0cq5lrjRNepAyfgM9vJ+9rUPDudpp5TLtA9qu9rGRB2xdw8iR
lv4uJZ17fdSeBpC2txP/bKtux70s6L1WbZKv1f8PxV5lOMdoQq7Vf8RAYTI/0CwlMVXw64d3ChgHxrNc
w4bppbU2yJ4zfblM1WctDMfqIX8Fn5aE36h+4ShymUq6oLfGR5RwKUVvChOIfN8/HyPspbo6H0+jaRT4
5+PL6eBk/HQy3Z8+6U83V/t/DM4vfxj8nQz+NRwcT8Pp4Gr/i38+3mw20/DLNnAbNsBNptHldP9/v59G
03C6mQ4+Xz0Lzqfn53a36f7k6cn3f8QpnPiDHQ6n302j6XS6uXoWBOdBtCXej9YUoBRCacXpS9ov5QP+
iuh42b7VbrUJGsJC9y1oH+yKaqv74CsXr/h6Osu1Fhz0XUYnnv3iQZwSpSbeTHOYaT5I6JzkqfbOLuyH
08gCnm3dRzteO2vB45TFN9+u4xbTQ4qOJKOiO6IfgDTsGFDHUCbZisg78/lWeUHrYrwWXKMpat+PNl7H
nNcH9yn4lqv0idxQBb+BR7wxeCOvD2EYOjtqbhpiUHCapGeniT4jp1Giz06T5Gx0GiXJWRjaW7YiN/Rz
QueMM5T455QpXTsAtJttuSPIA6JP0i0PhHfvht7hhW8jLK5m8iBK3fY8iW6J+obeNSGQypBkGeXJ6yVL
Ez/RLRxm24fC2CTZ2jZpbYvsXN7Qu6uvbZ7Ur1f7nBG8ecBzIVdEf9ZsRVVG6lciGfUhOWifScLmcyTn
AAaQjJpRwmeMETyvGmRz8M2CUxi2z8ICD7wW3xb9AP928VGEqD8RvQznqRDSbhDBq5eHQ/xXD1Fx5vtJ
OdWkdtmJ5vnLXVjcTBPJqhPJyx0oXm4jUJ0IRnb9thWCc/AT2Acv8QIYN4RNYH8Cvk/gyxdYBgi4RMCl
Awy6IVcGcoWQqwISI3kcUN6WU/iME2TLRizzFcFQiiRkllLIOdPWdJtPNaVKRVxXBFSQVMRwCqPhweEO
V4IA++D9yevQB5ybmP8jg+KbMYdavGG3NPEPkG3vr3/yTr5lm52IfioQ3ff37vf2ogg4WQNTQCBlWqOQ
NEuZvgMtQGkhKegl4wsTrLhEoA9KgF4SDRkVWUohJtyGmYybVfl8Hu6hbiDuMkGPIBUkcaQpg2xN0hyz
Xejd0LtescWvH965kH7PMkmS2lHd0Lv2UX3364d3HymR8fIXk4C0xZpJsZBUKd/7UUohxzCTYqOohERQ
hbmvyrNMSA0tPCG8/RGEhM2S6HOo+bmaiHmeprsNG6ebNlK/TG6UGQ0wGjFcNcLdKHLiN/9bcS3YmnK0
9Bj4J1Z6aofUzLKm2BpyQeNv0zWYPIrMlktz7qyN1ojbCk9ZtvoGxviITheAhHC6sbug7T3HeNvh0MKF
e0HTelcLvjP636C0Rc6SoSzuwixXy4+aaOrjmfXB/V9g2iLO3pAyrCyVaEWVIgtabBMLrkRKw1Qsyqm9
gjE0jk987w8JneULLzDOs+Blaxy5ggskMAhCLd6JmKT0E1vRQgZ4geELoHzcTjgy5ThAgpMGubZQhvwq
f05Yqvqg8jimChMmoUla0I9+DZN9+pZrCxn0jfkvB906R7auT1lMJyUipn4mP/tzNN54jmNw3lKVc6qa
UxVCO6erORcOuGvUO52dnWIcUITWKK9BQviCSu+sB/swh/3y+HqnEcKePeUzlZ2M7Z/t5Y4tu151rAe7
MrJ/EErDPs7OznpNYadisWZ0Y6LFPnCyKrXDmKf6AP7D76jpqVhY53G/tztF8RAzleFSr9LzOUvpBI+b
qphk1IanbscGRYK/YSl9x5RGzbF0UTR+BR2VTfwvoQHxmijMGbkoMjbFVvUUXCvBU8Zd3lNkQwCtm+Yh
mp/JimJEHkajF6+Onh+8eH5wNJiTIxo/j4+S4TF5Fc9nRy8Oh8dH86PjZP7qYDR6FeIOXr+JjTtMH+94
DJoqtNNM0zaY0kRqhNuaYKt3YmGyA0vJwavhYDic0Tk9fDV89XI2ovOjw/lsNjqeHc+OksPn9DAZKLbK
U6KFxBvdRpkRpajyxjBsTZiL0zGu2L+Qh8PnL1oTccoo17jk8qo1lVBbeGCCI/GflkxZzk0xjiqtYE0l
mzPjFogGh8v4YYXCMkUgSuIlCL2k0uQdbD6nEoP2lUioCqcc3mpjvBgiRMwbASqjMZszK27Vh5nQS8hN
IXBB9RKIdUSSGlwxdVv3DTxC6SVV1KAiiAuDrCmf8gH8bUkNKbjcLhrkPKFyYA62STluVV9jtm6CdKLp
Tbd1yDCCB1O7aC7jRgOtTWBojTRqr/nuLoGROVquyyvzHW+DKwWi0Q3CuZA/knjpl06WpnTVB8YTersV
o+Bcd7zX5RXF7Bom8JeP738OjbE1qINa0ccSZ5yaf2ngQ3MR+vjJGJi+HTSXwH52eftrd2ZmiKBDwE/F
3a0VJet7iNm1nbi3f0phbQsvvCCafMLPfsUuCm/sMFZnlJEFfWdKsWN4Ubs7JNfibyzB4TlJFa1mhEyo
HMPl5bAPPbwpvava/YlFmq+4GsNlQ85NqRunw3RK0bagzADTzJbmVDQ7A7M9re8yRJEQ3bV4Y8n3RkO6
6piWlBtGdlYaumLNIj4wsGWIsB0ilfrU3Lf19QGpdNjZmkCMYd7N8cHwe+/f3Pi1s4o7t453AhS7P9/a
3SgGOn5U7zF4NE1Zppj6nU7FmIVrwbjv9Vs5wn9yBr+YotsOOZjYdacIjv9zjTMlEmTMODY42y6UdEjC
e/qH24Oj0YsTOJ2dvSEsBR9jlRoeDGAjE8D6jVHzxTpXE+UGGGR5nfvd732dgue1jS1Wg9T73Q7n3VZo
ULcXNu54QEW7NNQYNjSaWyZv5wGizf2KchaRqVc2zjFYxnW/oyxIcv5vaurBcJjd/j7CeEib0dNjJAYT
F12b+o9VQRwPOlfNNIcJ5h6PqPGrVb0y7jWzlVSQhEoPpMDoXWmic+UBkYwMlixJKJ94WubUO3OJRzNX
udXeGcoYfJOoIBv70AsK2KKZ0OtWA7IxFaIOLbAX0Hn9PniXGINfecFD6oTs7oNnMyJEIsnmGxTJeer7
oob4xO9VgQPomUjuekEouN8z/YBeH4pGRr+7pxxFcOGK16YEAGQmcg0qM48A1Ni0xNQ4ihZUz4TQSkuS
mb5YImIVHYaHUVw0ZVVULGs2zXD0dWqtyAQ8BzSYGR2F5teBWnmNxSaDmlTRUijFxn/i6yVTAUZ2uK/f
07IXBCHC1r04rp+7XLFxWE0Qp54Yhjm0UmS+lzCFGyZeH1C5gi2eOJUwcavDeMnSRFLue6HT1iAkSWL4
9msSaKJJyYymnUhQa11ZwxgIxhdec63gF4Ija7XGJl4Mk6eu1KJ9l90eW8wZ+9CKfRx7oaQrsabbTDSh
0dfZvbvsh+GxxglN4P1fvY5gq0GgMY5ev1ngLq1bRvknqrQJs34hC+qbY25B3gNNFX0ERZh3oSOlyePJ
ei1WWa5Rf8mdAi7Gxik7ybfo6Cw8k6TkwFLfdydalDFtPeqJXzSXglBSktz5W/fYNusko2tq0jrTXhNz
c2dUq1ZRslvUK8IwdHr1JCTX5Nb3UlvzMBWF1OvX5OcKPuNmcaSyTijDXDaf8ZQa2VYNklKpfZyopW+F
eSsbEH8mPEkpcLJmC2ugjGFzHf4nvhdysh6kjN94xu55RUe0SinXzcrROjSVXf83D3Mnbwx0HWoiF1SH
LLmvCd+6WfOchiTJj2vKNXJMOZW+l4kMeUNVqGi7YCqrNcK3J9A4Gdtt2watWVAbpuMlVaCFOUhNZqBo
SmNNE5jdFVXqsKpSdW3huDVWyhQ9JoZr1DjfswUgR6C5tzjQfg7x/uL9GObsFpgGJWBDQS3FBgrlcdah
b4ro9oihZ692D/ycW5rRvgQgeEyB6Z4yKk+TcMcVMHTUjk3cIFHbDwr2qtr7osmaOc4AvnwBbz3IWJqq
wVKs6ECTmWvZ4DKa0pXLuG2ZfEFr0jCzT58aqFCTWV0wxZjvoTC8kiDbBdoyScXZ1I6zrMDhydZOcduc
Gah6dFp2Mxg3T8mqp4g1jd6rldmMapeCcGGLkUV1X502uKoCbtdvaD8KqQtDUBeD479xmuZbveViikhm
BncJ64XfLSUwdEAiOH3DK/aNL2cpLUq/ZSC2Xy4qz7gG15zctoNFNF9zH2si4fqft0t07k/CBdV/+fj+
Z7/+jnYfvKi+sL87iLZs+CaAqOyc4IbfC4xXbGm5RcN9YDK7ytbf9mGHLa14MuVpmFMdL9usmbyxdFFB
mzybIkBljJ3ziZ7twTP4hcq5kCuF2DSV5lGjbZl5P97SODcH+cHIxzNxlrn6e/Asqs6Z3tLYLq8/v8Lw
FWeccHs2hLOVryAkGfMD15Oyz5USSTa+Iy2CT67AbrZTpl8q1lSSNHX40DzkaXqCt+e9nSneQxvo2tZ1
0IqpOrBzbipfmac7Nfg3dgbcVLOt8UakCZU+OlMlchnTviu7FjJwiUEFUFRIQy+4HF65c4A35j2HKRGX
6OdCghQbSKg2ZYC6wO37D7xaF3bWT2q25B/JP4DZCyokWzBOUhtpi9k1je0LYlOnFpvyTmjzYq93mrB1
kVq5jQczceud9Vw36lbD/gS80+zsdHaG0fZpNDs7ncnoDNXPpo+1h+R+Etr2zj54p1F25rnUxhRQwlop
H76bmEcMlerXtqp198vR2dlFtboiog3oEtr6Y0G/i8oaKUHQsV1UUnFfY8Aphb0eYXFSD7Fi6DZghuZM
0rNTJKGDx8eS3klENxNmp9PI7NrmqIJK2LqYcvqrb/VJq4HWtnGVitc00Wid0jKPdW4e7jsPaYof9W6Z
SwRrTo4lzX5R2e+6oOvsIIOEqRgNwh2sD7vbX4/pFMVLGt8ocyNmRLEYYsHxcrknHqbLV98ok0KLWKQm
OMJVSqwo3HCx4aBonEtcuKHkhlNMpcI6OdYy/DeVigmuWi05D+l5TWz7rNU3HG0NVRIabdd+qs5gJsT8
I+FM371GPv314XD0POgoKLVF9QMos8zKByWxIjdoA82DF2J/ksCp3gh5A4rqPAP8aAWpcFvMaDv2KbuR
B8OD4WB4ODg4+DQ6Go8OxqPn4XD06vB4ODoc/b1rKeVJ98KjcPTyePRi9PL5cefCxv3oFGXZvPTGphrQ
74ZwVwupgCn/arm0duhv+Vzs3pkcvpxROjveCVEddwXb3w1ZnH9GUB0/pxim6ocWMK404ZoRTZMfdhzP
YXhw9OLwePTq4OjvD+FCr8jS7f0j148sGOjoHzfw/I2otzWy3Mk8suZ9v6vEdt/scxZ26iPVtmALjM9F
GR6Xge1nI1JXsTHFplo414TEm+QFxkT7j7TgiK/hf2o5i5lsWo5WPvfBVHGArjJ9B2sLA/YXIbVfCXU9
tH0Ab9mX7YCzj5U6VDWhKXXGvWvJ1+smTUkWzayGMLseJHcKqYh122UiE1oyhQHSkkrqTH98syEyGcRi
lRHNZvaNn/ndgUgTqJKbmkQr0WvER+aaSvhLzikcDA9G4eOY8qpMs9DF14KvqdRodJG5Mu+Jiaq328vf
7tyUx1m6kPrhmGW2Sd0Eurxxvez7Vva0ENrkEnal+1kUphi41Ax6ZQmnitAhJrynYUZBUjwckrJ/0aQP
G3QUNEF2Eqq0FKbSsULt0kt6B/S2YNKk6LWg/TsbhTey9Go6jFNKpB+EDmu9NIwir4HWhN2S9VvOdPmU
xyU2RUC8rDXw24nM7ha+kU93B380/LYW/sE3tPAxUWFSaTdvnqxy6NHbjPCkN7BVzq90rGpN4F4R+8eC
aynS3r/Tg3qow+V+9OFeyo+h13uoC7bVBGt5WmPBY8ETx/4YkI1HNug+dTvHxzXzj46//ypp8NHk9mPA
CAOEBC70499f6PyBfn8ztHlEG7Cx4KHudjOpQMof3+F+vqM1fbL3Ld1x0+M+eUyQVRzGi8coyqclxczd
hqrF0zBJMUhJ6rW8x7e6v/oew0R/jzicCvqhTq0hvnQD7X+ln7cxXUzfXqB/+DrqAn25js8FTGoLLyuM
V7sxiIVzNp1N1Try0EWK/eaoCayCb3/bYHZ+9DuTxhOkh+5wuydbhA+KagV5ZtTFWlmamOARii7GXu33
6WocRZvNJqzcZcipjugtWWUpVRHJWCTF5rMzu8ZdddXQupvAOgm37PXuHxlrWWuGxqlQVGnbY23+xkWK
DUxKN2g6s42fWqOVkGJj25ohUx+XYsP9oK1fhcAQHXql1PS5TFUcBmD2B9a8bRXWJUto++WWbjYuewr3
rRPf2RmMIniPW2pHS/eG/nZhDeds6zkIQtyrg56yFdxBTKPm3qxJFrKtVVpcTQSl67tH/lVxNWiyVKu6
boN2vZqscDrr3okWVc7Nd5Vst65RUcPtRF0WdZvUPIv27vf+LwAA//8kFrUDw0IAAA==
`,
	},

	"/assets.go": {
		name:    "assets.go",
		local:   "assets/assets.go",
		size:    0,
		modtime: 1612794766,
		compressed: `
H4sIAAAAAAAC/wEAAP//AAAAAAAAAAA=
`,
	},

	"/details_close.png": {
		name:    "details_close.png",
		local:   "assets/details_close.png",
		size:    686,
		modtime: 1603364414,
		compressed: `
H4sIAAAAAAAC/wCuAlH9iVBORw0KGgoAAAANSUhEUgAAABQAAAAUCAYAAACNiR0NAAACdUlEQVR4Aa2V
30tTYRjHVSgsoQKhXPUPlKZMlImiIIg/pggiiDeCBAdqyC4Kf9zVhVdeelX33QnCbtYvxZsxxkoqKxiz
1SzdFudMb5xO5/bt+fa+sPCw3IVf+PA+532e5ytn53lfK0qoSrggXBauCbUKxtxjjjVnq1IXXxFuCXeE
FqGdMOYec6zRtZX/M6sWrgv1hmHcDwQCLxOJxM9CoZAnjLnHHGt0bbXdVG1cEm4KLp/P9yKXyx2jhJhj
DWvZc9qUwUWhzuFwdMVisc8QFQ4PcbCygvT0NJLDw0gNDSH96DEyfj8KmQwo1rJHem8oD2VapX/sxlAo
9Aqi/N4e9hYWsDMwgARxD2DH7ZZYren5eZxYFqhgMOhnr3BVeSnn216v98Hf3yqXg/X0CX719YFs9/Wq
uJfPGtkzZ2fBWvZ4PB6DHtqrokaol7/0GqLM2hq2enpsxMmpvf3VVVD8UOJxV42Uet1Wy7J+M5mam8P3
7u6ySM7MgDJNM8mRUl5qYDtOREz+GB3FZlcXNjs7UUrMk/jYGKhsNnuo57TWZhgTw0hHO0EpRdpV/tvI
CKgjUdFQv/Lu7q7JZHxqCl9dLqGtSJtL4VJ80Wvc8xD2V9YfRUbmLZPp5WVstLTYaSWt+lmt1tISKH5Q
/VFqbGOTl6GNTk7iQ3Oz4FSrU8UfuTqdBJHxceQz+2APe/8Zm+Jgh8PhNxBlt7YQNQysNzVivbEJ7wXG
iiZEJiZwEI2C4mFgr/aoKnn0jk0TiWfPsTE4iHcN9xBuaMCn/n5sLy7iKJk8ffTqikfvfC4H9laWdX2l
UqltaDEu9/o664LtIOVfsOf4L+APb5yaiwyN8+8AAAAASUVORK5CYIIBAAD//1c8bPquAgAA
`,
	},

	"/details_open.png": {
		name:    "details_open.png",
		local:   "assets/details_open.png",
		size:    709,
		modtime: 1603364414,
		compressed: `
H4sIAAAAAAAC/wDFAjr9iVBORw0KGgoAAAANSUhEUgAAABQAAAAUCAYAAACNiR0NAAACjElEQVR4Aa2V
30uTURjHnVBUgglCaXVTJJVLMdhQm7+1JEvJiKAggmgSQiKJl0V0E13UVf0HQZmBLcvlpqa55tqV0YoM
zIhY6js0yDH11X17vuO8Sq5Zgl/48DznPD/G63nOMSmBkoUNwhYhTUgn9NUeY8z5p0wqOVXYKRwQLMJh
hYV7Kpaqck2rNdskbBPMdrv9osfjcQaDwW/RaHSR0OceY8xhrqox/a3ZZmGHkO9wOB7ouj6PBGKMOcxl
zcqmdDYKGZmZmSWjo6PvIYosRuDWXGj50IwTb4+hxleN5sAVvJh4jvBCGBRzWSO121UPk3EAaUKuz+d7
ycTp+Wnc/nwL1d4qHBVidmjZv/npBkJzGiiv19vFWmGrcVDsvKupqeky/056VMf1j9dQOViBKo9Aa/ix
dbnYcrQGWsBc1jQ2NtrZQ/VKShHM8kvdEPVN9qJsoBRl/aUoHyiJ+UpcCxJTtmfCDYoHJT2yYyOlPtca
CoUmGGwdbkVRbxGK+2y0gg2GisUvNmJCy/BVUJqm/eBIxXqpgbUtiBise12HQlc+ClwFSKRCiZH6wXpQ
s7OzETWn6XENa/trYemywuK0IpGsTotgRU3fcVBzouWG6pOnpqY0BhuGGpDXeUjIW7JKcfuX3tjjP9k4
FBkZN4PtY+042JGzzNMcGDJ35C7tmcU++tJmjE63OpSUuLEJ62GcfXUO+59kr8rJnlOY0WfAGtYaY/PH
YPv9fhdEY7++4nz/BWS17UPWY0Es2avWp3vOYOTnCCheBtaqHskJr95kZBL3AvdR2XkEex5mYbdQ+qwC
d97dRTAcXHn1Moyrt16PA2tN//V8jY+Pf4cS/VWerzU9sDay1gd2Xf4F/AZqlpeB9836LwAAAABJRU5E
rkJgggEAAP//wNSxe8UCAAA=
`,
	},

	"/index.html": {
		name:    "index.html",
		local:   "assets/index.html",
		size:    6301,
		modtime: 1612794322,
		compressed: `
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
`,
	},

	"/viewer.html": {
		name:    "viewer.html",
		local:   "assets/viewer.html",
		size:    2723,
		modtime: 1603658490,
		compressed: `
H4sIAAAAAAAC/6xW23LbyBF991e0R5Uqu0IA4k2RIICp2LGcOJZkO4pjPw6ABjHiYGZ2psGLt/bftwYA
Seiy+7C10gPJnr6cPt04g+TlP2/f3n3/9A4qquXiReI/QHK1TBkq5g3Ii8ULgKRG4pBX3DqklDVUBues
PSBBEhelkAhrgRu0SdSZDlGK15iyAl1uhSGhFYNcK0JFKfsHuJpL2YdCqS0QbmnvMAKhnBEWC8h2sBRU
NRk43dh8X409rsMbqrQdlKi0FDVXnaMUagUWZcoc7SS6CpEYVBbLlFVExsVR5IjnK8OpCjOtyZHlJi9U
mOs6OhiiWTgLx1Hu3NEW1kKFuXMMhCJcWkG7lLmKT89nwdeV1ufb2dv37nb6139V2/X5XfR5/ol/o9V/
qDlrlvO7W7z5erZ8I67w/5/eX91cXDfVbTmZfr74d3lfMcitdk5bsRQqZVxptat147qu2l78N4BQNXUc
Z1hqi/Bza/J/PRsxcCL7SgqFry/bw18OUUfnTNsCbWDFsqJ4bLbgtBQFnBTjYl7wy4PfRhRUxTDF+nJQ
SGobg11m/NXkb6PpeDSdj8Lpw2q5Lobg9BptKfUmhrVwIpN4TFdqRUHJayF3Mfz36lorHXzBZSO5Hb3V
ymnJ3eijyNByv1ngHUbXqKQe1VppZ3j+OJsTPzCG8cRsn6A+mcwmF5NBwEbbIthYbmJQ2tZcDo4qQRi0
BWIwFocNkg0rsaykJxCLQacZz1dLqxtVBPuSZVlmRTGMPqm5UMcYw4tCqOUQcu8nNS+yhkirQYl+KGf7
ofSIeCbxMe/Po7l8tAcx/M4K7FeFF6JxMUyHpPZIxqenfxlCKcQ69LKC9oik5nYpVAznZgun/v/ymfbP
hsl7c0DatGFPTzJNpOtHh8+1fFael/yPdE3aBBJLer79gVP7JD316p8GL3iB39bBaP6MvYeXojbaEld0
rJdEvVi0utEqMjibH9XP70h4/1ODdtdKXvc1mIbz8LRVuPsnAjeZnwXbmx83Ez6TtHozm13n0YcfU3M3
E/8b5zV+Ob1afXMzY5vow5Z/Tn9bzJKog7QYoKOdwZR5kqJ7vuadlXWgI25M0F0EHtez8aJIGW55bST6
JGyYsNSadbp5VwkHwoHTNYJrvdubKITvuoGcK/BPGyjcgL/rHKwFB6oQpCCSCEKZhoBneo0j4KoALp1u
M+8dGytDeBcuQ2B/9ynSrkpIW2JdCK5RQXtDkQbDLYncD9yb0IUv2mxSrPBpgpOP81PWTbfvP4m6qzvJ
dLFr6SjEuuXCq0vftLct+p1LSm1r0Mo1WS0oZUeu4xIpr66ExFevLy1SYxWUXDpki8O+ttlzyZ1LWUtF
4J8yM/AASDqOjvSzfYQvHfgrymrJwEieY6VlgTZlke+05agF73/5QbDoQebnqwfcGFTFAxAASS+ZHY6u
3S75UU8PyDJSkJEKjBU1t7tHqQCuPDWPbEM0PiVaBlZLX4w4NY4Bt4IHlSgKVCkj26Df3MEs+jxRB+ZB
pw/dHvxMIs9jP9njwRBOp7v9O9dBdw5dJc5w1S0JEmeLebd6ML64gDdJ5E8Prrz1s3wTNFayxRe+SSL+
tHh78ezLe23pqN6/vCVR69AuaBfVL2wSda+kvwYAAP//0Gc+kaMKAAA=
`,
	},

	"/": {
		name:  "/",
		local: `assets/`,
		isDir: true,
	},
}

var _escDirs = map[string][]os.FileInfo{

	"assets/": {
		_escData["/app-viewer.js"],
		_escData["/app.js"],
		_escData["/assets.go"],
		_escData["/details_close.png"],
		_escData["/details_open.png"],
		_escData["/index.html"],
		_escData["/viewer.html"],
	},
}
