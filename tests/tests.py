import unittest
import random
import string
from app.main import app


class BasicTests(unittest.TestCase):

    def setUp(self):
        app.config['TESTING'] = True
        app.config['WTF_CSRF_ENABLED'] = False
        app.config['DEBUG'] = False
        self.app = app.test_client()
        self.assertEqual(app.debug, False)

    def tearDown(self):
        pass

    def test_main_page(self):
        response = self.app.get('/',follow_redirects=True)
        self.assertEqual(response.status_code, 200)

    def test_peering_page(self):
        response = self.app.get('/peering.html', follow_redirects=True)
        self.assertEqual(response.status_code, 200)

    def test_404_page(self):
        response = self.app.get(f"/{''.join(random.choices(string.ascii_uppercase + string.digits, k=10))}",
                                follow_redirects=True)
        self.assertEqual(response.status_code, 404)

if __name__ == "__main__":
    unittest.main()