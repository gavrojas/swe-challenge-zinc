// composable, los demás componentes van a usar este store
import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { jwtDecode } from 'jwt-decode'
import type { JWTPayload } from '@/types/index'
import { useZincService } from '@/services/utils'

const { apiHandleEmails } = useZincService()

export const useAuthStore = defineStore('auth', () => {
  const router = useRouter()
  const session = ref('')
  const username = ref('')

  const isLoggedIn = computed(() => !!session.value ) 

  async function init() {
    const tokenStr = sessionStorage.getItem('token')
    if (tokenStr) {
      setSession(tokenStr, false) // false porque no quiero que se vuelva a redirigir al usuario
    }
  }

  function setSession(tokenStr: string, redirectToHome: boolean = true ) {
    const payload = jwtDecode(tokenStr) as JWTPayload
    console.log(payload);
    console.log(payload.MapClaims.eat);
    console.log("voy a imprimir eat ");
    console.log("payload mapclaims eat",payload.MapClaims.eat);
    console.log("terminé de imprimir eat ");
    // La diferencia entre sessionStorage y localStorage es que sessionStorage se limpia cuando se cierra el navegador, mientras que localStorage no, se va a quedar ahi poara siempre hasta que el sitio borre eso. 
    const now = new Date()
    const diff = payload.MapClaims.eat * 1000 - now.getTime()
    console.log(diff);

    // guardamos token en session storage para no perder el login mientras esté activo 
    sessionStorage.setItem('token', tokenStr)
    session.value = tokenStr

    setTimeout(()=>{
      clearSession()
    }, diff)

    if (redirectToHome) {
      router.push('/home')
    }

    router.push('/home')
  }

  async function clearSession(notifyLogout: boolean = true) { 
    if(notifyLogout){
      try {
        await apiHandleEmails('/logout', {notifyLogout: true})
      } catch (error) {
        console.error('Failed to logout:', error)
      }
    }

    session.value = ''
    sessionStorage.removeItem('token')
    router.push('/')
  }

  function isTokenExpired(): boolean {
    if (!session.value) return true
    const payload = jwtDecode(session.value) as JWTPayload
    console.log(payload);
    console.log(session.value);
    const now = new Date()
    return payload.MapClaims.eat * 1000 <= now.getTime()
  }

  function getToken() {
    if (isTokenExpired()) {
      clearSession()
      return ''
    }
    try {
      const payload = jwtDecode(session.value) as JWTPayload
      console.log(payload);
      console.log(session.value);
      if (!payload.MapClaims || !payload.MapClaims.eat) {
        clearSession()
        return ''
      }
      return session.value
    } catch (error) {
      clearSession()
      return ''
    }
  }

  return { isLoggedIn, init, setSession, clearSession, getToken, username }
})
