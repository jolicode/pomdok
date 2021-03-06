# This file was generated by GoReleaser. DO NOT EDIT.
class Pomdok < Formula
  desc "Simple wrapper to Symfony Go Binary for multi-app."
  homepage "https://github.com/jolicode/pomdok"
  version "1.2.0"
  bottle :unneeded

  if OS.mac?
    url "https://github.com/jolicode/pomdok/releases/download/v1.2.0/pomdok_1.2.0_Darwin_x86_64.tar.gz"
    sha256 "f41c4b3bb95cb8f158d7cac4acbfafb03bd145662373a741a18371e32aae8f2d"
  elsif OS.linux?
    if Hardware::CPU.intel?
      url "https://github.com/jolicode/pomdok/releases/download/v1.2.0/pomdok_1.2.0_Linux_x86_64.tar.gz"
      sha256 "214504f533471900fe019ed8afd925a9ee915ca90e55a8e7df448231af672d70"
    end
  end
  
  depends_on "wget"
  depends_on "nss"

  def install
    bin.install "pomdok"
  end
end
