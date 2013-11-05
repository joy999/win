//Author: 54zhua@gmail.com

package win

import (
	"syscall"
	"unsafe"
)

const (
	INTERNET_COOKIE_HTTPONLY      = 8192 //Requires IE 8 or higher
	INTERNET_COOKIE_THIRD_PARTY   = 131072
	INTERNET_FLAG_RESTRICTED_ZONE = 16
)

//const (
//	ERROR_INSUFFICIENT_BUFFER = 122
//	ERROR_INVALID_PARAMETER   = 87
//	ERROR_NO_MORE_ITEMS       = 259
//)

var (
	// Library
	libwininet uintptr

	// Functions
	internetGetCookieEx uintptr
)

func init() {
	//is64bit := unsafe.Sizeof(uintptr(0)) == 8

	// Library
	libwininet = MustLoadLibrary("wininet.dll")

	// Functions
	internetGetCookieEx = MustGetProcAddress(libwininet, "InternetGetCookieExW")
}

/*
BOOL InternetGetCookieEx(
  _In_         LPCTSTR lpszURL,
  _In_         LPCTSTR lpszCookieName,
  _Inout_opt_  LPTSTR lpszCookieData,
  _Inout_      LPDWORD lpdwSize,
  _In_         DWORD dwFlags,
  _In_         LPVOID lpReserved
);
lpszURL [in]
	A pointer to a null-terminated string that contains the URL with which the cookie to retrieve is associated. This parameter cannot be NULL or InternetGetCookieEx fails and returns an ERROR_INVALID_PARAMETER error.
lpszCookieName [in]
	A pointer to a null-terminated string that contains the name of the cookie to retrieve. This name is case-sensitive.
lpszCookieData [in, out, optional]
	A pointer to a buffer to receive the cookie data.
lpdwSize [in, out]
	A pointer to a DWORD variable.
	On entry, the variable must contain the size, in TCHARs, of the buffer pointed to by the pchCookieData parameter.
	On exit, if the function is successful, this variable contains the number of TCHARs of cookie data copied into the buffer. If NULL was passed as the lpszCookieData parameter, or if the function fails with an error of ERROR_INSUFFICIENT_BUFFER, the variable contains the size, in BYTEs, of buffer required to receive the cookie data.
	This parameter cannot be NULL or InternetGetCookieEx fails and returns an ERROR_INVALID_PARAMETER error.
dwFlags [in]
	A flag that controls how the function retrieves cookie data. This parameter can be one of the following values.
	Value	Meaning
INTERNET_COOKIE_HTTPONLY
	Enables the retrieval of cookies that are marked as "HTTPOnly".
	Do not use this flag if you expose a scriptable interface, because this has security implications. It is imperative that you use this flag only if you can guarantee that you will never expose the cookie to third-party code by way of an extensibility mechanism you provide.
	Version:  Requires Internet Explorer 8.0 or later.
INTERNET_COOKIE_THIRD_PARTY
	Retrieves only third-party cookies if policy explicitly allows all cookies for the specified URL to be retrieved.
INTERNET_FLAG_RESTRICTED_ZONE
	Retrieves only cookies that would be allowed if the specified URL were untrusted; that is, if it belonged to the URLZONE_UNTRUSTED zone.

lpReserved [in]
	Reserved for future use. Set to NULL.
*/
func InternetGetCookieEx(
	InLpszURL *uint16,
	InLpszCookieName *uint16,
	OutLpszCookieData *uint16,
	OutLpdwSize *uint,
	InDwFlags uint,
	InLpReserved uintptr) bool {
	ret, _, _ := syscall.Syscall6(internetGetCookieEx, 6,
		uintptr(unsafe.Pointer(InLpszURL)),
		uintptr(unsafe.Pointer(InLpszCookieName)),
		uintptr(unsafe.Pointer(OutLpszCookieData)),
		uintptr(unsafe.Pointer(OutLpdwSize)),
		uintptr(InDwFlags),
		InLpReserved)

	return BOOL(ret) == TRUE
}
