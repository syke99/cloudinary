# cloudinary
cloudinary aims to provide an improved experience when interacting with Cloudinary's REST API as compared to their provided Go SDK

Why was cloudinary built?
====
While working on a personal project, I came across the need for easily vectorizing raster/bitmap images. So I started looking around for packages to do the job, but didn't really find any. At least, not in Go.
So I pivoted to looking for third-party APIs and came across [Cloudinary](https://cloudinary.com/). While they provide a Go SDK, it seems a bit clunky and a little too verbose in my opinion. So I decided to
write a new tool to improve a developer's experience when working with Cloudinary's API.

What problem does cloudinary solve?
=====
cloudinary makes it simple for a developer to declare what kind of asset they're working with, apply any transformations to it they may wish, and upload the image (both with and without eager returning),
request an already uploaded image, transform an uploaded image that will remain on your Cloudinary account (won't be returned like when transforming and then requesting an image), and more. It makes
your code base cleaner, easier to read, and easier to maintain.



How do I use cloudinary?
====

### Installing
To install cloudinary, simply run

```bash
$ go get github.com/syke99/cloudinary
```

Then you can import the package in any go file you'd like

```go
import "github.com/syke99/cloudinary"
```

### Basic usage

Connect to the Cloudinary API and perform actions on assets

```go
// Initialize the Cloudinary instance by passing in an *http.Client, your Cloudinary cloud name,
// API Key, and API Secret
// *DON'T FORGET TO KEEP YOUR API KEY AND API SECRET HIDDEN (i.e. w/ environment variables)*
client := *http.Client
cloud := cloudinary.NewSignedCloudinary(client, <cloud_name>, <api_key>, <api_secret>)

// Call Image() or Video() to declare which type of asset you wish to work with, along with the name that
// you want your image to be saved as in your cloud account (no file extension needed)
img := cloud.Image("test_image")

// Create some transformations (they can be reused across multiple assets)
angle := cloud.NewAngle(45, resources.Ignore)
blue := cloud.NewColor("blue")

// Add some transformations to the asset
img.AddAngle(angle)
img.AddColor(blue)

// Call the type of action you want to perform on the image
// (RequestImage() requires the delivery type associated to the asset to be passed in)
response, err := img.RequestImage(resources.Upload)
```

You can also chain adding the transformations to an asset:

```go
img.AddAngle(angle).AddColor(blue)

response, err := img.RequestImage(resources.Upload)
```

You can even take it a step further and chain adding all your transformations and action together:

```go
response, err := img.AddAngle(angle).AddColor(blue).RequestImage(resources.Upload)
```


With this workflow, it's much simpler and quicker for developers to interact with Cloudinary's API by making reading, writing, and maintaining their interactions with the API at a glance a breeze. It also allows
for transformations, upload parameters, and more to be reused across multiple assets without having to either having to retype strings to be used in multiple places, or creating their own custom structs and
interfaces.


Who?
====

This library was developed by Quinn Millican (@syke99)


## License

This repo is under the MIT license, see [LICENSE](LICENSE) for details.