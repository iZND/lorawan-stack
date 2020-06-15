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

/* eslint-disable jest/valid-expect */

describe('Console logout', () => {
  it('Logs out successfully', () => {
    cy.login()
    cy.visit('/console')
    cy.grab('header__profile_dropdown').click()
    cy.findAllByText('Logout').click()

    cy.grab('login__form__user_name').should('exist')
    cy.url().should('include', '/login')
  })

  it('Obtains new CSRF token to logout if necessary', () => {
    cy.server()
    cy.route({
      method: 'POST',
      url: 'http://localhost:8080/console/api/auth/logout',
      onRequest: req => {
        expect(req.request.headers).to.have.property('X-CSRF-Token')
      },
    })

    cy.login()
    cy.visit('/console')
    cy.clearCookie('_console_csrf')
    cy.grab('header__profile_dropdown').click()
    cy.findAllByText('Logout').click()

    cy.grab('login__form__user_name').should('exist')
    cy.url().should('include', '/login')
  })
})
