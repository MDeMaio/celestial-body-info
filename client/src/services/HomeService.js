import http from "../http-common";

class HomeService {
  getAPOD() {
    return http.get(`/apod`);
  }

}

export default new HomeService();