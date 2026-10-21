---
title: "Installing CC65"
url: "/wiki/Installing_CC65/"
description: "An assembler is a program that translates assembly language source code into machine code. A commonly used assembler that produces machine code for 6502 CPUs is CA65, which is distributed as part of…"
wikiPage: "Installing_CC65"
---

An **assembler** is a program that translates assembly language source code into machine code. A commonly used assembler that produces machine code for 6502 CPUs is **CA65**, which is distributed as part of the **CC65** package. These instructions tell how to install and run CA65.

## Installing CC65 on Windows {#Installing_CC65_on_Windows}

### Downloading CC65 {#Downloading_CC65}

1.  Go to the links portion of the [main page of CC65's web site](https://cc65.github.io/index.html#Links).
2.  Click on the "Windows Snapshot" link. This will take you to CC65's SourceForge download page for `cc65-snapshot-win32.zip`.
    1.  There is also a 64-bit version available for computers with the appropriate architecture. Go to [CC65's SourceForge page's files section](https://sourceforge.net/projects/cc65/files/) and download `cc65-snapshot-win64.zip`.
    2.  Documentation is contained in `cc65-snapshot-docs.zip`.
3.  Extract `cc65-snapshot-win**.zip` to a known folder, such as `C:\cc65` for example.

### Adding CC65 to Path {#Adding_CC65_to_Path}

CC65 binaries are now contained in `<known_folder>\cc65\bin`. In order to run the binaries from anywhere, we must add the binary folder to the path environment variable.

1.  In the search bar, look for "Edit the system environment variables" and select it. Alternatively, type `sysdm.cpl` in the search bar and run it. You may need to confirm a UAC prompt.
2.  In the Advanced tab, click on the "Environment Variables..." button. A new window will appear.
3.  In the list of "System variables", select the "Path" variable and click the "Edit" button. A new window will appear with the list of paths.
4.  Click on the "New" button, and type the location of the CC65 binaries (`<known_folder>\cc65\bin`).
5.  Press OK on all three of the windows.
6.  You may need to restart any command-line terminals running, or log out and log back in.

### Showing file extensions {#Showing_file_extensions}

Windows is shipped with file name extensions hidden in the File Explorer. However, hiding them makes it easier to accidentally create a file name with two extensions, which may interfere with compiling. Therefore the first thing to do is turn on the display of file name extensions.

For Windows 8 and onward, this setting can be changed from the View tab in the File Explorer's ribbon interface. (Check the "Show File Extensions" checkbox).

For Windows 7 and earlier:

1.  Open Explorer and click on the Tools tab \> Folder Options...
2.  Go to the the View tab. In the scrolling list of Advanced Options, make sure that "Hide extensions for known file types" is *not* checked.
3.  Press OK to put the change into effect.

## Installing CC65 on Linux {#Installing_CC65_on_Linux}

### Installing with Package Managers {#Installing_with_Package_Managers}

The easiest method is to install CC65 using the package manager of your choice. Here are some examples from popular distributions:

- Debian-based: apt/apt-get
- Arch-based: yay (or any other [AUR helper](https://wiki.archlinux.org/title/AUR_helpers))
- RedHat-based: dnf/yum

### Building from git {#Building_from_git}

If CC65 is not available via your package manager, or you wish to use features/bugfixes not yet present in release versions, you can build it from the source code instead. On Debian-based distributions, open a terminal and enter the following commands. On other distributions, the `apt` command will need to be changed.

    sudo apt install build-essential git
    mkdir -p ~/develop
    cd ~/develop
    git clone https://github.com/cc65/cc65.git
    cd cc65
    nice make -j2
    make install PREFIX=~/.local
    which cc65

If your account has been configured to run applications built from source and installed for one user, the last step should show `/home/<username>/.local/bin/cc65`. If it does not, add `~/.local/bin` to your `PATH` environment variable:

    nano ~/.bashrc

    # and add the following at the end of the file
    if [ -d "$HOME/.local/bin" ] ; then
        PATH="$HOME/.local/bin:$PATH"
    fi

Press Ctrl+O Enter to save, then Ctrl+X to quit, and the change to `PATH` will take effect the next time you log in.

## Installing CC65 on Mac OS X {#Installing_CC65_on_Mac_OS_X}

Using Homebrew: On computer with Homebrew installed, open Terminal and type the following command: `brew install cc65`. Everything else should be automatic.

Using Macports: Alternatively macports can be used for easy and fast installation. Open a terminal and enter the following command: `sudo port install cc65`. This will install cc65, ca65 and ld65 on the computer.
