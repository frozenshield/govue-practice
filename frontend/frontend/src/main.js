import { createApp } from 'vue'
import App from './App.vue'
import PrimeVue from 'primevue/config';
import { msalInstance} from './msalConfig'

import 'primevue/resources/themes/lara-light-indigo/theme.css';
import 'primevue/resources/primevue.min.css';
import 'primeicons/primeicons.css';

//primevue imports
import InputText from 'primevue/inputtext';
import Button from 'primevue/button';
import Card from 'primevue/card';


async function bootstrap() {

    await msalInstance.initialize();

    msalInstance.handleRedirectPromise().then((response) => {
        if(response){
            msalInstance.setActiveAccount(response.account)


            console.log('Access token', response.accessToken)

            } else {
            const current = msalInstance.getActiveAccount()
            if (!current) {
                console.log('No active account')
            }
        }
    });
}





const app = createApp(App);

app.use(PrimeVue);
app.mount('#app')
app.component('InputText', InputText);
app.component('Button', Button);
app.component('Card', Card);


bootstrap()