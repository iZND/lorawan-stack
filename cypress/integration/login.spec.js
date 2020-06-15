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

describe('Login and OAuth authorization', () => {
  it('Displays an error on invalid login credentials', () => {
    cy.visit('/console')

    cy.grab('login__form__user_name')
      .type('adminwrong')
      .should('have.value', 'adminwrong')
    cy.grab('login__form__password')
      .type('adminwrong')
      .should('have.value', 'adminwrong')
      .type('{enter}')

    cy.location('pathname').should('include', '/oauth')

    cy.findAllByText('incorrect password or user ID').should('exist')
  })
  it('Logs into the console successfully', () => {
    cy.visit('/console')

    cy.grab('login__form__user_name')
      .type('admin')
      .should('have.value', 'admin')
    cy.grab('login__form__password')
      .type('admin')
      .should('have.value', 'admin')
      .type('{enter}')

    cy.grab('overview__welcome_text', { timeout: 10000 }).should('contain.text', 'Welcome')
  })
  it('Logs in', () => {
    cy.login()
    cy.visit('/console')
  })
})
