$(document).ready(function () {
    let lastRequestTime = 0;
    const RATE_LIMIT_MS = 1000; // 1 second

    $('#shortenForm').on('submit', function (e) {
        e.preventDefault();

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

        // Show loading state
        setLoading(true);

        // Make API request
        $.ajax({
            url: '/api/links',
            method: 'POST',
            contentType: 'application/json',
            data: JSON.stringify({ targetUrl }),
            success: function (response) {
                lastRequestTime = Date.now();
                const shortUrl = response.url;
                const validUntil = new Date(response.validUntil);

                $('#targetUrl').attr('readonly', true);
                $('#shortenBtn').attr('disabled', true).hide();

                $('#shortUrl').val(shortUrl);
                $('#validUntil').text(validUntil.toLocaleString());

                $('#resultCard').fadeIn();
            },
            error: function (xhr) {
                let errorMessage = 'An error occurred while shortening the URL';
                if (xhr.responseJSON && xhr.responseJSON.error) {
                    errorMessage = xhr.responseJSON.error;
                }
                showError(errorMessage);
            },
            complete: function () {
                setLoading(false);
            }
        });
    });

    $('#copyBtn').on('click', function () {
        const shortUrl = $('#shortUrl').val();
        navigator.clipboard.writeText(shortUrl).then(function () {
            const originalText = $('#copyBtn').text();
            $('#copyBtn').text('Copied!');
            setTimeout(() => $('#copyBtn').text(originalText), 2000);
        });
    });

    // Show loading state
    function setLoading(loading) {
        const $btn = $('#shortenBtn');
        const $spinner = $btn.find('.spinner-border');
        const $text = $btn.find('.btn-text');

        if (loading) {
            $btn.prop('disabled', true);
            $spinner.show();
            $text.text('Shortening...');
        } else {
            $btn.prop('disabled', false);
            $spinner.hide();
            $text.text('Shorten');
        }
    }

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
