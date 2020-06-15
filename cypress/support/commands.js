// Copyright © 2020 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

Cypress.Commands.add('login', () => {
  const baseUrl = Cypress.config('baseUrl')
  // Obtain csrf token
  cy.request({
    method: 'GET',
    url: `${baseUrl}/oauth/login`,
  })

  cy.getCookie('_oauth_csrf').then(c => {
    // Login to OAuth provider.
    cy.request({
      method: 'POST',
      url: `${baseUrl}/oauth/api/auth/login`,
      body: { user_id: 'admin', password: 'admin' },
      headers: {
        'X-CSRF-Token': c.value,
      },
    })
  })

  // Do OAuth round trip
  cy.request({
    method: 'GET',
    url: `${baseUrl}/console/login/ttn-stack?next=/`,
  })

  // Obtain access token
  cy.request({
    method: 'GET',
    url: `${baseUrl}/console/api/auth/token`,
  }).then(resp => {
    window.localStorage.setItem('accessToken-1259217480', JSON.stringify(resp.body))
  })
})

Cypress.Commands.add('grab', (selector, options) => cy.get(`[data-test="${selector}"]`, options))

Cypress.Commands.add('seedDatabase', () => {
  cy.exec('node .mage/restore-db-dump.js')
    .its('code')
    .should('eq', 0)
})
