cmake_minimum_required(VERSION 4.0.3)
project(dywoqcxx CXX)

set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED True)

file(GLOB SOURCES src/*.cxx)

add_executable(${PROJECT_NAME} ${SOURCES})
target_link_options(${PROJECT_NAME} PRIVATE c)
target_include_directories(${PROJECT_NAME} PUBLIC include/)
target_compile_features(${PROJECT_NAME} PRIVATE -fno-exceptions -fno-rtti)
target_compile_definitions(${PROJECT_NAME} PUBLIC -D__dywoqcxx=202507)