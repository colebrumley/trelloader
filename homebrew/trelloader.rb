require 'rbconfig'
class Trelloader < Formula
  desc "Create a Trello board from a JSON template."
  homepage "https://github.com/colebrumley/trelloader"
  version "0.1.0"

  if Hardware::CPU.is_64_bit?
    case RbConfig::CONFIG['host_os']
    when /mswin|msys|mingw|cygwin|bccwin|wince|emc/
      :windows
    when /darwin|mac os/
      url "https://github.com/colebrumley/trelloader/releases/download/v0.1.0/trelloader_0.1.0_darwin_amd64.zip"
      sha256 "97b6f153e89f5caa9761653f17ec5c02298b3a10813c0993da44e1d0d611ea2a"
    when /linux/
      url "https://github.com/colebrumley/trelloader/releases/download/v0.1.0/trelloader_0.1.0_linux_amd64.tar.gz"
      sha256 "7a2f28b986a165308ae8f2139c89e5ab8629f40ba131452e4815db53b410e56c"
    when /solaris|bsd/
      :unix
    else
      :unknown
    end
  else
    case RbConfig::CONFIG['host_os']
    when /mswin|msys|mingw|cygwin|bccwin|wince|emc/
      :windows
    when /darwin|mac os/
      url "https://github.com/colebrumley/trelloader/releases/download/v0.1.0/trelloader_0.1.0_darwin_386.zip"
      sha256 "5bbe9998a861dbdc28c7864aa48d3cc940b86231161095dc2da5d8b162943c10"
    when /linux/
      url "https://github.com/colebrumley/trelloader/releases/download/v0.1.0/trelloader_0.1.0_linux_386.tar.gz"
      sha256 "a03d391551b467144717905290e68cf873cf726cf04286813302398b2548dcd5"
    when /solaris|bsd/
      :unix
    else
      :unknown
    end
  end

  def install
    bin.install "trelloader"
  end

  test do
    system "trelloader"
  end

end
