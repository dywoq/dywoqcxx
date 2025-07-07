#ifndef _DYWOQCXX___CONFIG_HXX
#define _DYWOQCXX___CONFIG_HXX

#ifdef __dywoqcxx
#    define _DYWOQCXX_VERSION __dywoqcxx
#else
#    define _DYWOQCXX_VERSION 202507
#endif

#if defined(_MSC_VER)
#    define _DYWOQC_EXPORT_FROM_ABI __declspec(dllexport)
#elif defined(__clang__) || defined(__GNUC__)
#    define _DYWOQC_EXPORT_FROM_ABI __attribute__((visibility("default")))
#else
#    define _DYWOQC_EXPORT_FROM_ABI
#endif

#if defined(__clang__) || defined(__GNUC__)
#    define _DYWOQC_HIDE_FROM_ABI __attribute__((visibility("hidden")))
#else
#    define _DYWOQC_HIDE_FROM_ABI
#endif

#endif