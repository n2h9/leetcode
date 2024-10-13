cmake_minimum_required(VERSION 3.25)

project(
    solution 
    VERSION 1.0.0
    LANGUAGES CXX
)


if(NOT CMAKE_BUILD_TYPE)
    set(CMAKE_BUILD_TYPE Release CACHE STRING "" FORCE)
endif()
message(STATUS "CMAKE_BUILD_TYPE: ${CMAKE_BUILD_TYPE}")

add_library(cxx_setup INTERFACE)
target_compile_options(cxx_setup INTERFACE -Wall -Wpedantic -Wextra)
target_compile_features(cxx_setup INTERFACE cxx_std_20)
target_include_directories(cxx_setup INTERFACE ${PROJECT_SOURCE_DIR})
message(STATUS "PROJECT_SOURCE_DIR: ${PROJECT_SOURCE_DIR}")


include(FetchContent)

# google tests
FetchContent_Declare(
  googletest
  URL https://github.com/google/googletest/releases/download/v1.15.2/googletest-1.15.2.tar.gz
)
# For Windows: Prevent overriding the parent project's compiler/linker settings
set(gtest_force_shared_crt ON CACHE BOOL "" FORCE)

# end of googletests


FetchContent_MakeAvailable(
    googletest
)

include(CTest)

if (BUILD_TESTING)
    message(STATUS "BUILD_TESTING is on")
    add_executable(solution_test solution_test.cpp)
    target_link_libraries(solution_test PRIVATE GTest::gtest_main)

    include(GoogleTest)
    gtest_discover_tests(solution_test)
endif()

add_library(solution solution.cpp)
target_link_libraries(solution PUBLIC cxx_setup)
