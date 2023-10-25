class MySpotifyCli < Formula
  desc "Go Spotify Cli"
  homepage "https://github.com/envoy49/go-spotify-cli"
  url "https://github.com/envoy49/go-spotify-cli/gsc/gsc.tar.gz"
  sha256 "281c283a7cecda6773b3fb9f9ac236e43ecdc3c085bd37cf3ff3f72b9f5fdd56"

  def install
    bin.install "gsc"
  end
end