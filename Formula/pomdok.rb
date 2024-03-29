# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Pomdok < Formula
  desc "Simple wrapper to Symfony Go Binary for multi-app."
  homepage "https://github.com/jolicode/pomdok"
  version "1.3.1"
  depends_on :macos

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/jolicode/pomdok/releases/download/v1.3.1/pomdok_1.3.1_Darwin_arm64.tar.gz"
      sha256 "df8da8b137bb725fc69feedba385bb67b04a76eaa2ddd67294b8839752c446a1"

      def install
        bin.install "pomdok"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/jolicode/pomdok/releases/download/v1.3.1/pomdok_1.3.1_Darwin_amd64.tar.gz"
      sha256 "21b7b57f08dbbcb0221a09978f899bf707ca696dc5454afd881dd16007d1bf20"

      def install
        bin.install "pomdok"
      end
    end
  end

  depends_on "wget"
  depends_on "nss"
end
