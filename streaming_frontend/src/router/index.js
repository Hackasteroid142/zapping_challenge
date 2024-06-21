
/**
 * router/index.ts
 *
 * Automatic routes for `./src/pages/*.vue`
 */

// Composables
import { createRouter, createWebHistory } from 'vue-router/auto'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
})

router.beforeEach((to, from, next) => {
  // Verificar si la ruta requiere autenticación y si el usuario está autenticado
  if (to.meta.requiresAuth && !(localStorage.getItem('token') !== null)) {
    // Redirigir a la página de inicio de sesión u otra página
    next('/');
  } else {
    // Permitir el acceso a la ruta
    next();
  }
});

export default router
