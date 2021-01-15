import axios from 'axios';

const API_URL = 'http://localhost:8090/api/auth';

class AuthService {
  login(user) {
    return axios({
        method: 'post',
        url: API_URL + '/login',
        headers: {"content-type": "text/plain" },
        data: {
          email: user.email,
          password: user.password
        }
      })
      .then(response => {
        console.log(response.data.email);
        if (response.data.token) {
          localStorage.setItem('user', JSON.stringify(response.data));
        }
        return response.data;
      });
  }

  logout() {
    localStorage.removeItem('user');
  }

  register(user) {
    console.log(user);
    return axios({
      method: 'post',
      url: API_URL + '/register',
      headers: {"content-type": "text/plain" },
      data: {
        name: user.name,
        email: user.email,
        password: user.password
      }
    })
    .then(response => {
      console.log(response);
    });
  }
}

export default new AuthService();
