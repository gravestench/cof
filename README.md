<!-- PROJECT LOGO -->
<h1 align="center">COF</h1>
<p align="center">
  COF file decoder.
   <br />
 An arcane format, used by Blizzard's Diablo II. 
  <br />
  <br />
  <a href="https://github.com/gravestench/cof/issues">Report Bug</a>
  ·
  <a href="https://github.com/gravestench/cof/issues">Request Feature</a>
</p>

<!-- ABOUT THE PROJECT -->
## About

The COF file codec is an essential library for handling composite animation
files within Blizzard's iconic game, Diablo II. These COF files are
binary-encoded and reside within the MPQ archives. They play a crucial role in
representing complex animations by combining multiple DCC or DC6 images,
creating dynamic visuals for in-game characters and objects. The animations are
often composed of various layers, such as a left arm holding a sword, a right
arm with a shield, and a head with a helmet. The COF files

Key Components of the COF Codec:
--------------------------------

1. COF Structure: The COF contains information about the number of directions,
   frames per direction, and the total number of layers involved in the animation.

2. CofLayer: holds data related to individual layers, including their types,
   shadows, transparency, and draw effects. Additionally, it can be associated with
   a specific weapon class.

3. FrameEvent: The COF codec defines FrameEvent as a single frame in an
   animation. These frame events are utilized to compose animations and bring the
   in-game characters and objects to life.

4. Animation Priority: The COF codec maintains an internal priority matrix that
   determines the order in which layers are composited during animation playback.
   This priority matrix ensures that elements are properly displayed on top of each
   other, preserving the correct visual hierarchy.

Codec Operations:
-----------------

The COF codec enables both marshaling and unmarshaling operations. Marshaling
takes a COF structure and converts it into a binary byte slice, while
unmarshaling does the opposite, converting a byte slice back into a COF
structure. These operations are crucial for reading COF files from the MPQ
archives, handling their content, and rendering the animations within the
game engine.

COF in Diablo II:
------------------

COF files play an integral role in animating characters, monsters, and various
game elements in Diablo II. By utilizing this specialized codec, the game
achieves smooth and immersive animations that contribute to the rich and
captivating gameplay experience.

Please note that the provided code represents a Golang implementation of the
COF codec, showcasing how it is utilized to parse and handle COF from the
vanilla game files.

This package provides a COF animation transcoder implementation for Diablo II.

It offers the necessary tools to read, write, and manipulate COF animation files, which are vital for creating captivating and dynamic animations for characters and objects within the game.

This package also includes command-line and graphical applications for interacting with COF files.

## Project Structure
* `pkg/` - This directory contains the core COF transcoder library. Import this package to write new Golang applications using the COF functionalities.
  ```golang
  import (
    "github.com/gravestench/cof"
  )
  ```
* `cmd/` - This directory contains command-line and graphical applications, each having their sub-directory.
* `assets/` - This directory stores various files, such as images displayed in this README or test COF file data.

## Getting Started

### Prerequisites
Ensure you have [Go 1.16][golang] or a later version installed, and your Go environment is correctly set up.
To use the applications found in `cmd/`, you need to define `$GOBIN` and add it to your `$PATH` environment variable.
```shell
export GOBIN=$HOME/.gobin
mkdir -p $GOBIN
PATH=$PATH:$GOBIN
```

### Installation
With `$GOBIN` defined and on your `$PATH`, you can build and install all apps inside `cmd/` using these commands:

```shell
# Clone the repository and navigate to the directory
git clone http://github.com/gravestench/cof
cd cof

# Build and install inside $GOBIN
go build ./cmd/...
go install ./cmd/...
```

After this, you should be able to run the apps inside `cmd/` from the command-line, like `cof-view`.

<!-- CONTRIBUTING -->
## Contributing

The COF repository follows a similar project structure to other transcoder libraries. The `~/pkg/` directory holds the core transcoder library, and `~/cmd/` contains subdirectories for each CLI/GUI application that can be compiled.

Contributions are **highly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<!-- MARKDOWN LINKS & IMAGES -->
[cof]: https://github.com/gravestench/cof
[dc6]: https://github.com/gravestench/dc6
[golang]: https://golang.org/dl/