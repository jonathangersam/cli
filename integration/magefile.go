// +build mage

// Creates the binary in the current directory.  It will overwrite any existing
// binary.
func GolangCrossBuild() error {

	return mage.GolangCrossBuild(mage.DefaultGolangCrossBuildArgs())

}

// Sends the binary to the server.
func Test() error {
	return nil
}

func BuildGoDaemon() error {

	return mage.BuildGoDaemon()

}

// CrossBuild cross-builds the beat for all target platforms.

func CrossBuild() error {

	return mage.CrossBuild()

}
