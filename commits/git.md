# github.com/git/git

## commit message

```sh
commit eb1c42da8e21cc2a8dacd21023a179b788858887 (HEAD)
Author: Jeff King <peff@peff.net>
Date:   Thu May 18 15:21:02 2023 -0400

    t/lib-httpd: make CGIPassAuth support conditional

    Commit 988aad99b4 (t5563: add tests for basic and anoymous HTTP access,
    2023-02-27) added tests that require Apache to support the CGIPassAuth
    directive, which was added in Apache 2.4.13. This is fairly old (~8
    years), but recent enough that we still encounter it in the wild (e.g.,
    RHEL/CentOS 7, which is not EOL until June 2024).

    We can live with skipping the new tests on such a platform. But
    unfortunately, since the directive is used unconditionally in our
    apache.conf, it means the web server fails to start entirely, and we
    cannot run other HTTP tests at all (e.g., the basic ones in t5551).

    We can fix that by making the config conditional, and only triggering it
    for t5563. That solves the problem for t5551 (which then ignores the
    directive entirely). For t5563, we'd see apache complain in start_httpd;
    with the default setting of GIT_TEST_HTTPD, we'd then skip the whole
    script.

    But that leaves one small problem: people may set GIT_TEST_HTTPD=1
    explicitly, which instructs the tests to fail (rather than skip) when we
    can't start the webserver (to avoid accidentally missing some tests).

    This could be worked around by having the user manually set
    GIT_SKIP_TESTS on a platform with an older Apache. But we can be a bit
    friendlier by doing the version check ourselves and setting an
    appropriate prereq. We'll use the (lack of) prereq to then skip the rest
    of t5563. In theory we could use the prereq to skip individual tests, but
    in practice this whole script depends on it.

    Reported-by: Todd Zullinger <tmz@pobox.com>
    Signed-off-by: Jeff King <peff@peff.net>
    Signed-off-by: Junio C Hamano <gitster@pobox.com>
```

## diff

```sh
diff --git a/t/lib-httpd.sh b/t/lib-httpd.sh
index 2c49569f67..7f6db7d078 100644
--- a/t/lib-httpd.sh
+++ b/t/lib-httpd.sh
@@ -181,6 +181,20 @@ enable_http2 () {
 	test_set_prereq HTTP2
 }
 
+enable_cgipassauth () {
+	# We are looking for 2.4.13 or more recent. Since we only support
+	# 2.4 and up, no need to check for older major/minor.
+	if test "$HTTPD_VERSION_MAJOR" = 2 &&
+	   test "$HTTPD_VERSION_MINOR" = 4 &&
+	   test "$(echo $HTTPD_VERSION | cut -d. -f3)" -lt 13
+	then
+		echo >&4 "apache $HTTPD_VERSION too old for CGIPassAuth"
+		return
+	fi
+	HTTPD_PARA="$HTTPD_PARA -DUSE_CGIPASSAUTH"
+	test_set_prereq CGIPASSAUTH
+}
+
 start_httpd() {
 	prepare_httpd >&3 2>&4
 
diff --git a/t/lib-httpd/apache.conf b/t/lib-httpd/apache.conf
index 76335cdb24..4eef1d1266 100644
--- a/t/lib-httpd/apache.conf
+++ b/t/lib-httpd/apache.conf
@@ -138,7 +138,9 @@ Alias /auth/dumb/ www/auth/dumb/
 <LocationMatch /custom_auth/>
 	SetEnv GIT_EXEC_PATH ${GIT_EXEC_PATH}
 	SetEnv GIT_HTTP_EXPORT_ALL
+	<IfDefine USE_CGIPASSAUTH>
 	CGIPassAuth on
+	</IfDefine>
 </LocationMatch>
 ScriptAlias /smart/incomplete_length/git-upload-pack incomplete-length-upload-pack-v2-http.sh/
 ScriptAlias /smart/incomplete_body/git-upload-pack incomplete-body-upload-pack-v2-http.sh/
diff --git a/t/t5563-simple-http-auth.sh b/t/t5563-simple-http-auth.sh
index ccf7e54b07..1ec51e0a2b 100755
--- a/t/t5563-simple-http-auth.sh
+++ b/t/t5563-simple-http-auth.sh
@@ -5,6 +5,12 @@ test_description='test http auth header and credential helper interop'
 . ./test-lib.sh
 . "$TEST_DIRECTORY"/lib-httpd.sh
 
+enable_cgipassauth
+if ! test_have_prereq CGIPASSAUTH
+then
+	skip_all="no CGIPassAuth support"
+	test_done
+fi
 start_httpd
 
 test_expect_success 'setup_credential_helper' '
```
