$(document).ready(function () {
    let lastRequestTime = 0;
    const RATE_LIMIT_MS = 1000; // 1 second

    $('#shortenForm').on('submit', function (e) {
        const now = Date.now();
        if (now - lastRequestTime < RATE_LIMIT_MS) {
            showError('Please wait a second before shortening another link');
            return false;
        }

        const targetUrl = $('#targetUrl').val();
        if (!isValidUrl(targetUrl)) {
            showError('Please enter a valid HTTPS URL');
            return false;
        }

        // Reset alerts
        $('.alert').hide();

        return true;
    });

    $('#copyBtn').on('click', function () {
        const shortUrl = $('#shortUrl').val();
        navigator.clipboard.writeText(shortUrl).then(function () {
            const originalText = $('#copyBtn').text();
            $('#copyBtn').text('Copied!');
            setTimeout(() => $('#copyBtn').text(originalText), 2000);
        });
    });

    function showError(message) {
        $('#errorAlert').text(message).fadeIn();
    }

    function isValidUrl(url) {
        try {
            const urlObj = new URL(url);
            if (urlObj.protocol !== 'https:') {
                return false;
            }
            return true;
        } catch {
            return false;
        }
    }
});
