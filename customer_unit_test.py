import requests

"""
Checks that the server returns a valid response
"""
def test(host, port):
	response = requests.get("http://" + host + ":" + str(port) + "/customer?customer=392")
	assert response.status_code == 200, "Status code was {} instead of 200".format(response.status_code)
	result = response.json()
	assert result["Name"] == "Trom Chocolatier"
	assert result["Location"] == "577,322"

if __name__ == '__main__':

	test("hotrod-customer", 8081)
