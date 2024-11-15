Feature: Reset password with 2FA

Scenario: User reset password with 2FA
  When User send "POST" request to "/reset_password"
  Then the response on /reset_password code should be 200
  And the response on /reset_password should match json:
      """
      {
        "message": "Password reset code sent to email"
      }
      """
  And user send "POST" request to "/verify_reset_password"
  Then the response on /verify_reset_password code should be 200
  And the response on /verify_reset_password should match json:
      """
      {
        "message": "Password changed successfully!"
      }
      """