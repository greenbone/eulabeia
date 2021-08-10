# see /opt/local/share/cmake-3.16/Modules/FindGTest.cmake
find_package(cgreen 1.1.0)
# Now simply link against gtest as needed. Eg
add_executable(all_tests src/all.c src/start_scan.c src/scan_porgress.c)
target_link_libraries(all_tests ${CGREEN_LIBRARIES} Eulabeia::eulabeia)


add_test(NAME test_start_scan_fail
    COMMAND all_tests start_scan_returns_error_on_publish_fail)
add_test(NAME test_start_scan_success
    COMMAND all_tests start_scan_success)
add_test(NAME test_progress_fail
    COMMAND all_tests scan_progress_failures)
