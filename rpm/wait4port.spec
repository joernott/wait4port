Name:          wait4port
Version:       %
Release:       1
Epoch:         0
Summary:       Wait for tcp port to open
AutoReqProv:   no
BuildRoot:     %buildroot
BuildArch:     x86_64
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

%clean
rm -rf $RPM_BUILD_ROOT/*

%files
%defattr(-,root,root,-)
/usr/bin/wait4port
