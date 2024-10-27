cmake_minimum_required(VERSION 3.14)
project(prefix_sum_test)

# 设置编译器标志
# GoogleTest requires at least C++14
set(CMAKE_CXX_STANDARD 14)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

# 通过cmake的模块，设置 Google Test
include(FetchContent)
FetchContent_Declare(
  googletest
  URL https://github.com/google/googletest/archive/refs/tags/v1.15.2.zip
)
# For Windows: Prevent overriding the parent project's compiler/linker settings
set(gtest_force_shared_crt ON CACHE BOOL "" FORCE)
FetchContent_MakeAvailable(googletest)

# 添加子目录，包含 Google Test
enable_testing()

# 添加单元测试
add_executable(
	prefix_sum_test
	prefix_sum_test.cpp
	prefix_sum.cpp
)
# 添加依赖关系
target_link_libraries(prefix_sum_test gtest gtest_main)

include(GoogleTest)
gtest_discover_tests(prefix_sum_test)

