;nsis of gdb

Name 'gdb' ; name of application

OutFile 'gdbInstaller.exe' ; name of installer application

RequestExecutionLevel admin  ; executable authority

Unicode True

InstallDir $PROGRAMFILES\GDB  ; default install directory

InstallDirRegKey HKLM 'Software\GDB' 'Install_Dir'

Icon "gdb.ico"  ; icon of installer

ShowInstDetails show

ShowUninstDetails show

VIAddVersionKey  "ProductName" "gdbInstaller"
VIAddVersionKey  "Comments" "gdbInstaller Programmer"
VIAddVersionKey  "CompanyName" "Southeast University"
VIAddVersionKey  "LegalCopyright" "Copyright (C) 2020. KeepSilence. All Rights Reserved."
VIAddVersionKey  "FileDescription" "gdbInstaller Programmer"
VIAddVersionKey  "FileVersion" "1.0.0"
VIAddVersionKey "ProductVersion" "1.0.0"

VIProductVersion "1.0.0.0"

VIFileVersion "1.0.0.0"


; Pages

Page components 
Page directory
Page instfiles 

UninstPage uninstConfirm
UninstPage instfiles


; the stuff to install

Section "gdb"
    SectionIn RO
    SetOutPath $INSTDIR
    File /r ssl ; files to install
    File gdb.exe
    File gdbCli.exe
    File config.json

    WriteRegStr HKLM SOFTWARE\GDB "GDB_Install" "$INSTDIR"  ; Write the installation path into the registry
     ; Write the uninstall keys for Windows
    WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\GDB" "DisplayName"  "gdb"
    WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\GDB" "UninstallString" '"$INSTDIR\uninstall.exe"'
    WriteRegDWORD HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\GDB" "NoModify" 1
    WriteRegDWORD HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\GDB" "NoRepair" 1
    WriteUninstaller "$INSTDIR\uninstall.exe"
SectionEnd

Section "Start Menu Shortcuts"
    CreateDirectory "$SMPROGRAMS\GDB"
    CreateShortcut "$SMPROGRAMS\GDB\Uninstall.lnk" "$INSTDIR\uninstall.exe"
    CreateShortcut "$SMPROGRAMS\GDB\gdb.lnk" "$INSTDIR\gdb.exe" 
SectionEnd

Section "Desktop ShortCut" 
    CreateShortcut "$DESKTOP\gdb.lnk" "$INSTDIR\gdb.exe" 
SectionEnd

Section "Boot"
    CreateShortcut "$SMPROGRAMS\Startup\gdb.lnk" "$INSTDIR\gdb.exe" 
SectionEnd

Section "Uninstall"
  
  ; Remove registry keys
  DeleteRegKey HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\GDB"
  DeleteRegKey HKLM SOFTWARE\GDB

  ; Remove files and uninstaller
  Delete $INSTDIR\uninstall.exe

  ; Remove shortcuts, if any
  Delete "$SMPROGRAMS\GDB\*.lnk"

  ; Remove directories
  RMDir /r "$SMPROGRAMS\GDB"
  RMDir /r "$INSTDIR"

SectionEnd
