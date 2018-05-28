Name:          wait4port
Version:       0.1.0
Release:       1
Epoch:         0
Summary:       Wait for tcp port to open
AutoReqProv:   no
BuildRoot:     %buildroot
BuildArch:     x86_64
Source0:       wait4port
Prefix:        /
Group:         Application/Internet
License:       MIT
URL:           https://github.com/joernott/wait4port
Packager:      Joern Ott
Provides:      wait4port = %{version}-%{release}

%description
wait4port is a simple commandline tool to wait for a tcp port to become open.

%prep
%build
%install
cd $RPM_BUILD_ROOT
mkdir -p $RPM_BUILD_ROOT/usr/bin
cp %{SOURCE0} $RPM_BUILD_ROOT/usr/bin/

%clean
rm -rf $RPM_BUILD_ROOT/*

%files
%defattr(-,root,root,-)
/usr/bin/wait4port
