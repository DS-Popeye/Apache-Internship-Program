# Licensed to the Apache Software Foundation (ASF) under one
# or more contributor license agreements.  See the NOTICE file
# distributed with this work for additional information
# regarding copyright ownership.  The ASF licenses this file
# to you under the Apache License, Version 2.0 (the
# "License"); you may not use this file except in compliance
# with the License.  You may obtain a copy of the License at
# http://www.apache.org/licenses/LICENSE-2.0
# Unless required by applicable law or agreed to in writing,
# software distributed under the License is distributed on an
# "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
# KIND, either express or implied.  See the License for the
# specific language governing permissions and limitations
# under the License.

from setuptools import setup, find_packages
from age import VERSION 

with open("README.md", "r", encoding='utf8') as fh:
    long_description = fh.read()

setup(
    name             = 'apache-age-python',
    version          = VERSION.VERSION,
    description      = 'Python driver support for Apache AGE',
    long_description=long_description,
    long_description_content_type="text/markdown",
    author           = 'MD RABIUL AW',
    author_email     = 'rabiul35-487@diu.edu.bd',
    url              = 'https://github.com/DS-Popeye/Apache-Internship-Program/tree/main/Question-3',
    download_url     = 'https://github.com/DS-Popeye/Apache-Internship-Program/tree/main/Question-3' ,
    license          = 'Apache2.0',
    install_requires = [ 'psycopg2', 'antlr4-python3-runtime==4.9.2' ],
    packages         = ['age', 'age.gen'],
    keywords         = ['Graph Database', 'Apache AGE', 'PostgreSQL'],
    python_requires  = '>=3.9',
    # package_data     =  {},
    # zip_safe=False,
    classifiers      = [
        'Programming Language :: Python :: 3.9'
    ]
)
