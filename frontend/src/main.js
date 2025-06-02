import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import PrimeVue from 'primevue/config';
import { msalInstance} from './msalConfig'


import 'primevue/resources/themes/lara-light-indigo/theme.css';
import 'primevue/resources/primevue.min.css';
import 'primeicons/primeicons.css';
import ConfirmationService from 'primevue/confirmationservice';

//primevue imports
import InputText from 'primevue/inputtext';
import Button from 'primevue/button';
import Card from 'primevue/card';
import Dialog from 'primevue/dialog';
import Message from 'primevue/message';
import ToastService from 'primevue/toastservice';
import ConfirmDialog from 'primevue/confirmdialog';



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
app.use(router)
app.use(ToastService)
app.use(ConfirmationService); 
app.mount('#app')
app.component('ConfirmDialog', ConfirmDialog)
app.component('InputText', InputText);
app.component('Button', Button);
app.component('Card', Card);
app.component('Dialog', Dialog)
app.component('Message', Message)
app.component('Toast', ToastService)


bootstrap()