set BiosPassword=InsertPasswordHere
cctk\cctk.exe ^
  --SetupPwd= ^
  --ValSetupPwd=%BiosPassword% ^
  --UsbPortsFront=Enabled ^
  --UsbPortsRear=Enabled ^
  --EmbNic1=EnabledPxe ^
  --UefiNwStack=Enabled ^
  --SmartErrors=Enabled ^
  --SecureBootMode=DeployedMode ^
  --Virtualization=Enabled ^
  --WakeOnLan=LanOnly ^
  --AcPwrRcvry=On